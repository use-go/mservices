package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/micro/micro/v3/client/cli/util"
	"github.com/micro/micro/v3/cmd"
	"github.com/micro/micro/v3/service/client"
	"github.com/micro/micro/v3/service/context"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(&cli.Command{
		Name:  "network",
		Usage: "Manage the micro service network",
		Subcommands: []*cli.Command{
			{
				Name:   "connect",
				Usage:  "connect to the network. specify nodes e.g connect ip:port",
				Action: util.Print(networkConnect),
			},
			{
				Name:   "connections",
				Usage:  "List the immediate connections to the network",
				Action: util.Print(networkConnections),
			},
			{
				Name:   "graph",
				Usage:  "Get the network graph",
				Action: util.Print(networkGraph),
			},
			{
				Name:   "nodes",
				Usage:  "List nodes in the network",
				Action: util.Print(networkNodes),
			},
			{
				Name:   "routes",
				Usage:  "List network routes",
				Action: util.Print(networkRoutes),
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "service",
						Usage: "Filter by service",
					},
					&cli.StringFlag{
						Name:  "address",
						Usage: "Filter by address",
					},
					&cli.StringFlag{
						Name:  "gateway",
						Usage: "Filter by gateway",
					},
					&cli.StringFlag{
						Name:  "router",
						Usage: "Filter by router",
					},
					&cli.StringFlag{
						Name:  "network",
						Usage: "Filter by network",
					},
				},
			},
			{
				Name:   "services",
				Usage:  "Get the network services",
				Action: util.Print(networkServices),
			},
		},
	})
}

func networkConnect(c *cli.Context, args []string) ([]byte, error) {
	if len(args) == 0 {
		return nil, nil
	}

	request := map[string]interface{}{
		"nodes": []interface{}{
			map[string]interface{}{
				"address": args[0],
			},
		},
	}

	var rsp map[string]interface{}

	req := client.DefaultClient.NewRequest("network", "Network.Connect", request, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	b, _ := json.MarshalIndent(rsp, "", "\t")
	return b, nil
}

func networkConnections(c *cli.Context, args []string) ([]byte, error) {

	request := map[string]interface{}{
		"depth": 1,
	}

	var rsp map[string]interface{}

	req := client.DefaultClient.NewRequest("network", "Network.Graph", request, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	if rsp["root"] == nil {
		return nil, nil
	}

	peers := rsp["root"].(map[string]interface{})["peers"]

	if peers == nil {
		return nil, nil
	}

	b := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(b)
	table.SetHeader([]string{"NODE", "ADDRESS"})

	// root node
	for _, n := range peers.([]interface{}) {
		node := n.(map[string]interface{})["node"].(map[string]interface{})
		strEntry := []string{
			fmt.Sprintf("%s", node["id"]),
			fmt.Sprintf("%s", node["address"]),
		}
		table.Append(strEntry)
	}

	// render table into b
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()

	return b.Bytes(), nil
}

func networkGraph(c *cli.Context, args []string) ([]byte, error) {

	var rsp map[string]interface{}

	req := client.DefaultClient.NewRequest("network", "Network.Graph", map[string]interface{}{}, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	b, _ := json.MarshalIndent(rsp, "", "\t")
	return b, nil
}

func networkNodes(c *cli.Context, args []string) ([]byte, error) {

	var rsp map[string]interface{}

	// TODO: change to list nodes
	req := client.DefaultClient.NewRequest("network", "Network.Nodes", map[string]interface{}{}, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	// return if nil
	if rsp["nodes"] == nil {
		return nil, nil
	}

	b := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(b)
	table.SetHeader([]string{"ID", "ADDRESS"})

	// get nodes

	if rsp["nodes"] != nil {
		// root node
		for _, n := range rsp["nodes"].([]interface{}) {
			node := n.(map[string]interface{})
			strEntry := []string{
				fmt.Sprintf("%s", node["id"]),
				fmt.Sprintf("%s", node["address"]),
			}
			table.Append(strEntry)
		}
	}

	// render table into b
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()

	return b.Bytes(), nil
}

func networkRoutes(c *cli.Context, args []string) ([]byte, error) {

	query := map[string]string{}

	for _, filter := range []string{"service", "address", "gateway", "router", "network"} {
		if v := c.String(filter); len(v) > 0 {
			query[filter] = v
		}
	}

	request := map[string]interface{}{
		"query": query,
	}

	var rsp map[string]interface{}

	req := client.DefaultClient.NewRequest("network", "Network.Routes", request, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	if len(rsp) == 0 {
		return []byte(``), nil
	}

	b := bytes.NewBuffer(nil)
	table := tablewriter.NewWriter(b)
	table.SetHeader([]string{"SERVICE", "ADDRESS", "GATEWAY", "ROUTER", "NETWORK", "METRIC", "LINK"})

	routes := rsp["routes"].([]interface{})

	val := func(v interface{}) string {
		if v == nil {
			return ""
		}
		return v.(string)
	}

	var sortedRoutes [][]string

	for _, r := range routes {
		route := r.(map[string]interface{})
		service := route["service"]
		address := route["address"]
		gateway := val(route["gateway"])
		router := route["router"]
		network := route["network"]
		link := route["link"]
		metric := route["metric"]

		var metInt int64
		if metric != nil {
			metInt, _ = strconv.ParseInt(route["metric"].(string), 10, 64)
		}

		// set max int64 metric to infinity
		if metInt == math.MaxInt64 {
			metric = "∞"
		} else {
			metric = fmt.Sprintf("%d", metInt)
		}

		sortedRoutes = append(sortedRoutes, []string{
			fmt.Sprintf("%s", service),
			fmt.Sprintf("%s", address),
			fmt.Sprintf("%s", gateway),
			fmt.Sprintf("%s", router),
			fmt.Sprintf("%s", network),
			fmt.Sprintf("%s", metric),
			fmt.Sprintf("%s", link),
		})
	}

	sort.Slice(sortedRoutes, func(i, j int) bool { return sortedRoutes[i][0] < sortedRoutes[j][0] })

	table.AppendBulk(sortedRoutes)
	// render table into b
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.Render()

	return b.Bytes(), nil
}

func networkServices(c *cli.Context, args []string) ([]byte, error) {

	var rsp map[string]interface{}

	req := client.DefaultClient.NewRequest("network", "Network.Services", map[string]interface{}{}, client.WithContentType("application/json"))
	err := client.DefaultClient.Call(context.DefaultContext, req, &rsp, client.WithAuthToken())
	if err != nil {
		return nil, err
	}

	if len(rsp) == 0 || rsp["services"] == nil {
		return []byte(``), nil
	}

	rspSrv := rsp["services"].([]interface{})

	var services []string

	for _, service := range rspSrv {
		services = append(services, service.(string))
	}

	sort.Strings(services)

	return []byte(strings.Join(services, "\n")), nil
}
