/* pam_local.c */
#define PAM_SM_AUTH

#include <security/pam_modules.h>
#include <security/pam_ext.h>
#include <syslog.h>
#include <stddef.h>
#include <string.h>
#include <stdlib.h>

/* Authenticate entry point */
PAM_EXTERN int
pam_sm_authenticate(pam_handle_t *pamh, int flags, int argc, const char **argv)
{
    const char *user = NULL;
    int pam_rc;

    /* get the username */
    pam_rc = pam_get_item(pamh, PAM_USER, (const void**)&user);
    if (pam_rc != PAM_SUCCESS || user == NULL) {
        pam_syslog(pamh, LOG_ERR, "pam_local: no username provided");
        return PAM_USER_UNKNOWN;
    }
    pam_syslog(pamh, LOG_INFO, "pam_local: attempting auth for user %s", user);

    /* prompt for password */
    char *password = NULL;
    pam_rc = pam_prompt(pamh, PAM_PROMPT_ECHO_OFF, &password, "Password: ");
    if (pam_rc != PAM_SUCCESS || password == NULL) {
        pam_syslog(pamh, LOG_ERR, "pam_local: prompt failed");
        return PAM_AUTH_ERR;
    }
    pam_syslog(pamh, LOG_DEBUG, "pam_local: got password of length %zu", strlen(password));

    /* free the response buffer */
    free(password);

    return PAM_SUCCESS;
}

/* Set credentials (post-auth) */
PAM_EXTERN int
pam_sm_setcred(pam_handle_t *pamh, int flags, int argc, const char **argv)
{
    return PAM_SUCCESS;
}
