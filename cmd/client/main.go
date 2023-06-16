package main

import (
	"context"
	"fmt"
	"strings"

	"wormhole/client"
	"wormhole/rpc/gen-go/common"

	"github.com/apache/thrift/lib/go/thrift"
	tea "github.com/charmbracelet/bubbletea"
)

var defaultCtx = context.Background()

var Database string = "default"

func handleClient(client *common.RPCServiceClient, sql string) (*common.QueryResp, error) {
	client.Ping(defaultCtx)

	req := &common.QueryReq{
		QueryID:   2,
		Statement: sql,
		Database:  Database,
	}

	resp, _ := client.ExecuteQueryStatement(context.Background(), req)
	return resp, nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, cfg *thrift.TConfiguration, sql string) *common.QueryResp {
	var transport thrift.TTransport
	if secure {
		transport = thrift.NewTSSLSocketConf(addr, cfg)
	} else {
		transport = thrift.NewTSocketConf(addr, cfg)
	}
	transport, _ = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
	}
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	resp, _ := handleClient(common.NewRPCServiceClient(thrift.NewTStandardClient(iprot, oprot)), sql)
	return resp
}

func rpcExec(sql string) *common.QueryResp {
	// var protocolFactory thrift.TProtocolFactory
	// switch *protocol {
	// case "compact":
	protocolFactory := thrift.NewTCompactProtocolFactoryConf(nil)

	// var transportFactory thrift.TTransportFactory
	transportFactory := thrift.NewTTransportFactory()

	addr := "10.67.0.15:1320"
	secure := false
	var cfg thrift.TConfiguration

	// if err := runClient(transportFactory, protocolFactory, addr, secure, &cfg); err != nil {
	// 	fmt.Println("error running server:", err)
	// }
	resp := runClient(transportFactory, protocolFactory, addr, secure, &cfg, sql)
	if strings.HasPrefix(sql, "use") {
		Database = string(resp.Output)
	}
	return resp
}

// // A simple program demonstrating the text input component from the Bubbles
// // component library.

type model struct {
	input *client.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// By default, the prompt component will not return a "tea.Quit"
	// message unless Ctrl+C is pressed.
	//
	// If there is no error in the input, the prompt component returns
	// a "common.DONE" message when the Enter key is pressed.
	switch msg {
	case client.DONE:
		return m, tea.Quit
	}

	_, cmd := m.input.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.input.View()
}

func (m model) Value() string {
	return m.input.Value()
}

func main() {
	for {
		m := model{input: &client.Model{ValidateFunc: client.VFNotBlank}}
		p := tea.NewProgram(&m)
		p.Run()
		if m.input.Canceled() {
			break
		}
		resp := rpcExec(m.Value())
		fmt.Println(resp.QueryID)
		fmt.Println(string(resp.Output))
	}

}
