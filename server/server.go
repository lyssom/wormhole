package server

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"context"
	"crypto/tls"
	"fmt"

	"wormhole/rpc/gen-go/common"
	"wormhole/server/query"
	"wormhole/server/query/interpreter"
	"wormhole/server/query/parser"
	"wormhole/server/query/statement"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/apache/thrift/lib/go/thrift"
)

type RPCServiceHandler struct {
}

func NewRPCServiceHandler() *RPCServiceHandler {
	return &RPCServiceHandler{}
}

func (p *RPCServiceHandler) Ping(ctx context.Context) (err error) {
	fmt.Print("ping()\n")
	return nil
}

type SqlParserVisitor struct {
	*parser.BaseSqlParserVisitor
}

func (p *RPCServiceHandler) ExecuteQueryStatement(ctx context.Context, req *common.QueryReq) (_r *common.QueryResp, _err error) {
	state := ExecuteParser(req.Statement)
	resp := common.NewQueryResp()
	resp.QueryID = req.QueryID
	resp.Output = Execute(state.(statement.Statement), req.Database)

	return resp, nil
}

func ExecuteParser(sql string) interface{} {
	is := antlr.NewInputStream(sql)
	lexer := parser.NewSqlLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	par := parser.NewSqlParser(stream)
	visitor := query.NewVisitor()
	statement := par.Statement().Accept(visitor)
	return statement
}

func Execute(stat statement.Statement, usedDB string) []byte {

	var resp []byte
	interp := interpreter.Interpreter{UsedDB: usedDB}
	switch stat.GetKind() {
	case "select":
		resp = interp.ExecuteFetchColumns(stat.(statement.QueryStatement))
	case "create database":
		dbName := stat.(statement.CreateDatabaseStatement).Database
		resp = interp.ExecuteCreateDatabase(dbName)
	case "create table":
		resp = interp.ExecuteCreateTable(stat.(statement.CreateTableStatement))
	case "insert":
		resp = interp.ExecuteInsert(stat.(statement.InsertStatement))
	case "use":
		resp = interp.ExecuteUse(stat.(statement.UseStatement))
	}

	return resp
}

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	rpc := NewRPCServiceHandler()
	handler := common.NewRPCServiceProcessor(rpc)
	server := thrift.NewTSimpleServer4(handler, transport, transportFactory, protocolFactory)

	fmt.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
