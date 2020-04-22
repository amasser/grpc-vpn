package cmd

import (
	"github.com/fatih/color"
	"github.com/gjbae1212/grpc-vpn/pkg/vpn-server/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
)

var (
	runCmd = &cobra.Command{
		Use:    "run",
		Short:  "Start vpn-server",
		Long:   "Start vpn-server",
		PreRun: startPreRun(),
		Run:    startRun(),
	}
)

func startPreRun() commandRun {
	return func(cmd *cobra.Command, args []string) {}
}

func startRun() commandRun {
	return func(cmd *cobra.Command, args []string) {
		// apply default params
		var opts []server.Option
		if defaultConfig.SubNet != "" {
			opts = append(opts, server.WithVpnSubNet(defaultConfig.SubNet))
		}
		if defaultConfig.Port != "" {
			opts = append(opts, server.WithGrpcPort(defaultConfig.Port))
		}
		if defaultConfig.JwtSalt != "" {
			opts = append(opts, server.WithVpnJwtSalt(defaultConfig.JwtSalt))
		}
		if defaultConfig.TlsCertification != "" {
			opts = append(opts, server.WithGrpcTlsCertification(defaultConfig.TlsCertification))
		}
		if defaultConfig.TlsPem != "" {
			opts = append(opts, server.WithGrpcTlsPem(defaultConfig.TlsPem))
		}

		// apply auth interceptors
		var interceptors []grpc.UnaryServerInterceptor
		if auth1, ok := defaultConfig.Auth.AuthForGoogleOpenID(); ok {
			interceptors = append(interceptors, auth1)
		}
		if auth2, ok := defaultConfig.Auth.AuthForAwsIAM(); ok {
			interceptors = append(interceptors, auth2)
		}
		if len(interceptors) > 0 {
			opts = append(opts, server.WithGrpcUnaryInterceptors(interceptors))
		}

		// create server
		server, err := server.NewVpnServer(opts...)
		if err != nil {
			log.Panicln(color.RedString("[ERR] %s", err.Error()))
		}
		if err := server.Run(); err != nil {
			log.Panicln(color.RedString("[ERR] %s", err.Error()))
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
