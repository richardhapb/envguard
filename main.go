package main

import (
	"envguard/env"
	"errors"
	"fmt"
	"os"
	"slices"
)

func main() {
	runCheck(os.Args)
}

// Execute the `check` subcommand, parse the variables from each .env
// and compare each pair --env .env-file, printing the results
func runCheck(args []string) {
	usage := "Usage: envguard check --env .env"

	if len(args) < 2 {
		fmt.Println(usage)
		fmt.Println("If --env is omitted, the default is to check .env")
		os.Exit(1)
	}

	switch args[1] {
	case "check":
		env_file := ".env"
		if len(args) > 2 && !slices.Contains(args, "--env") {
			fmt.Printf("Incorrect param: %s\n", args[len(args)-1])
			fmt.Println(usage)
			os.Exit(1)
		}

		if len(args) > 4 {
			envs, err := matchEnvs(args[2:])
			if err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			envs_instances := make([]*env.Env, len(envs))

			for i, env_elem := range envs {
				envs_instances[i] = env.New(env_elem)
			}

			var prev_instance *env.Env

			if len(envs_instances) > 1 {
				for i, env_instance := range envs_instances {
					env_instance.Parse()
					if i%2 != 0 && i > 1 {
						unique1, unique2 := prev_instance.Compare(env_instance)
						fmt.Printf("\nUnique values in %s\n\n", prev_instance.Path)
						fmt.Println(unique1)
						fmt.Println()
						fmt.Println("===============================================")
						fmt.Printf("\nUnique values in %s\n\n", env_instance.Path)
						fmt.Println(unique2)
					}

					prev_instance = env_instance
				}

				return
			} else if len(args) == 4 {
				env_file = envs_instances[0].Path
			}
		}

		env_element := env.New(env_file)
		_, err := env_element.Parse()

		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println(env_element)

	default:
		fmt.Printf("Invalid command: %s\n", args[1])
		fmt.Println(usage)
		os.Exit(1)
	}
}

// Retrieve all environment file names from the command line arguments
func matchEnvs(args_pairs []string) ([]string, error) {
	error := errors.New("env parameters must be in pairs: --env env_file_path")

	if len(args_pairs)%2 != 0 {
		return nil, error
	}

	envs := make([]string, len(args_pairs)/2)

	for i, arg := range args_pairs {
		if i%2 == 0 {
			if arg != "--env" {
				return nil, error
			}
			continue
		}

		envs = append(envs, arg)
	}

	return envs, nil
}
