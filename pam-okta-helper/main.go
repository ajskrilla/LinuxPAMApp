package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "golang.org/x/term"
  "syscall"
  "pam-okta-helper/oktaauth"
  "pam-okta-helper/util"
  //db mod
  "pam-okta-helper/db"
)



func main() {
	if err := util.InitLogger("auth.log"); err != nil {
		fmt.Fprintf(os.Stderr, "failed to init logger: %v\n", err)
		os.Exit(1)
	}
  //init the db
  err := db.Init("pamcache.db")
  if err != nil {
    log.Fatalf("DB Init error: %v", err)
  }

	// Make oktaauth use our global app logger
	oktaauth.SetLogger(util.Logger)

	util.Logger.Println("Starting authentication flow")


	// io.MultiWriter will write to all writers you give it.
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Username: ")
  user, _ := reader.ReadString('\n')
  user = strings.TrimSpace(user)

  fmt.Print("Password: ")
  pwBytes, _ := term.ReadPassword(int(syscall.Stdin))
  password := string(pwBytes)
  fmt.Println()

  resp, err := oktaauth.Login(user, password)
  if err != nil {
    util.Logger.Printf("Login error: %v\n", err)
    return
  }
  util.Logger.Println("Status:", resp.Status)

  // 2) Handle MFA_REQUIRED
  if resp.Status == "MFA_REQUIRED" {
    util.Logger.Println("Available factors:")
    for i, factor := range resp.Embedded.Factors {
      util.Logger.Printf("%d) ID=%s type=%s (profile: %+v)\n",
        i+1, factor.ID, factor.FactorType, factor.Profile)
    }

    util.Logger.Printf("Pick factor number: ")
    var choice int
    fmt.Scanln(&choice)
    factorID := resp.Embedded.Factors[choice-1].ID

    util.Logger.Printf("Enter passCode: ")
    passCode, _ := reader.ReadString('\n')
    passCode = strings.TrimSpace(passCode)

    // 3) Call VerifyFactor()
    verfResp, err := oktaauth.VerifyFactor(factorID, passCode)
    if err != nil {
      util.Logger.Println("Verify error:", err)
      return
    }
    util.Logger.Println("Verify Status:", verfResp.Status)
    if verfResp.Status == "SUCCESS" {
      util.Logger.Println("Got sessionToken:", verfResp.SessionToken)
    }
  } else if resp.Status == "SUCCESS" {
    util.Logger.Println("Got sessionToken:", resp.SessionToken)
  } else {
    util.Logger.Println("Unhandled status:", resp.Status)
  }
}
