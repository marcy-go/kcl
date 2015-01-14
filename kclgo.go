package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
  "strings"
  "flag"
  "path/filepath"
  "github.com/mitchellh/cli"
)

func pkgDest(a, v string) string {
  return fmt.Sprintf("%s-%s.jar", a, v)
}

func pkgUrl(a, v, d string) string {
  b := "http://search.maven.org/remotecontent?filepath="
  return fmt.Sprintf("%s%s/%s/%s/%s", b, strings.Join(strings.Split(a, "."), "/"), v, d, pkgDest(v, d))
}

func jarPath() string {
  pwd, _  := os.Getwd()
  return pwd + separator() + "kcl_jars"
}

func separator() string {
  return string(os.PathSeparator)
}

func kclClassPath(prop string) (string, error) {
  p, err := filepath.Glob(jarPath() + separator() + "*")
  if err != nil {
    return "", err
  }
  cpath := strings.Join(p, ":")
  return cpath + ":" + filepath.Dir(prop), nil
}

func multiLangDaemonClass() string {
  return "com.amazonaws.services.kinesis.multilang.MultiLangDaemon"
}

type Setup struct{}

func (s *Setup) Help() string {
  return "kclgo setup has option."
}

func (s *Setup) Synopsis() string {
  return "Get Kinesis Client Library from Maven projects."
}

func (s *Setup) Run(args []string) int {

  MavenPackages := [][]string{
    {"com.amazonaws", "amazon-kinesis-client", "1.2.0"},
    {"com.fasterxml.jackson.core", "jackson-core", "2.1.1"},
    {"org.apache.httpcomponents", "httpclient", "4.2"},
    {"org.apache.httpcomponents", "httpcore", "4.2"},
    {"com.fasterxml.jackson.core", "jackson-annotations", "2.1.1"},
    {"commons-codec", "commons-codec", "1.3"},
    {"joda-time", "joda-time", "2.4"},
    {"com.amazonaws", "aws-java-sdk", "1.7.13"},
    {"com.fasterxml.jackson.core", "jackson-databind", "2.1.1"},
    {"commons-logging", "commons-logging", "1.1.1"},
  }

  if err := os.MkdirAll(jarPath(), 0755); err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
    return 1
  }

  for _, pkg := range MavenPackages {
    dest := jarPath() + separator() + pkgDest(pkg[1], pkg[2])
    url  := pkgUrl(pkg[0], pkg[1], pkg[2])

    response, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      return 1
    }

    fmt.Println("status:", response.Status)
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      return 1
    }

    fmt.Printf("Saving %s -> %s\n",url ,dest)
    file, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0755)
    if err != nil {
      fmt.Fprintf(os.Stderr, err.Error())
      return 1
    }

    file.Write(body)
    file.Close()
  }
  return 0
}

type Run struct{}

func (r *Run) Help() string {
  return "kclgo run -j <path-to-java> -p <path-to-properties>"
}

func (r *Run) Synopsis() string {
  return "Run Kinesis Client Library."
}

func (r *Run) Run(args []string) int {
  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, `
      Usage of %s:
      %s [OPTIONS] ARGS...
      Options\n`, os.Args[0],os.Args[0])
      flag.PrintDefaults()
  }
  f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
  var java, prop string
  f.StringVar(&java, "j", "java", "Path to java binary")
  f.StringVar(&java, "java", "java", "Path to java binary")
  f.StringVar(&prop, "p", "", "Path to properties file")
  f.StringVar(&prop, "properties", "", "Path to properties file")
  f.Parse(args)
  if prop == "" {
    fmt.Fprintf(os.Stderr, "-p option is required.\n")
    return 1
  }
  cpath, err := kclClassPath(prop)
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
    return 1
  }
  fmt.Println(java, "-cp", cpath, multiLangDaemonClass(), filepath.Base(prop))
  err = exec.Command(java, "-cp", cpath, multiLangDaemonClass(), filepath.Base(prop)).Run()
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
    return 1
  }
  return 0
}

func main() {
  c := cli.NewCLI("kclgo", "0.1.0")
  c.Args = os.Args[1:]
  c.Commands = map[string]cli.CommandFactory{
    "setup": func() (cli.Command, error) {
      return &Setup{}, nil
    },
    "run": func() (cli.Command, error) {
      return &Run{}, nil
    },
  }
  ret, err := c.Run()
  if err != nil {
    fmt.Fprintf(os.Stderr, err.Error())
  }
  os.Exit(ret)
}
