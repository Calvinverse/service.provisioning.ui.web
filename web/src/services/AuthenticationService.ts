import * as Msal from 'msal'
import msalConfig from '../config/msalConfig'
import MsalConfig from '../config/msalConfig'

class AuthenticationService {
  app: Msal.UserAgentApplication

  constructor() {
    // let redirectUri = window.location.origin;
    let redirectUri = msalConfig.redirectUrl.toString()
    let PostLogoutRedirectUri = '/'
    this.graphUrl = msalConfig.graphEndpoint.toString()
    this.applicationConfig = new ApplicationAuthenticationConfig{
      clientID: msalConfig.clientID,
      authority: config.authority,
      graphScopes: config.graphscopes
    }
    this.app = new Msal.UserAgentApplication(
      new Msal.Configuration(
      this.applicationConfig.clientID,
      this.applicationConfig.authority,
      () => {
        // callback for login redirect
      },
      {
        redirectUri
      }
      )
    )
  }

  // Core Functionality
  loginPopup() {
    return this.app.loginPopup(this.applicationConfig.graphScopes).then(
      idToken => {
        const user = this.app.getUser();
        if (user) {
          return user;
        } else {
          return null;
        }
      },
      () => {
        return null;
      }
    );
  }

  loginRedirect() {
    this.app.loginRedirect(this.applicationConfig.graphScopes)
  }

  logout() {
    this.app._user = null
    this.app.logout()
  }

  // Graph Related
  getGraphToken() {
    return this.app.acquireTokenSilent(this.applicationConfig.graphScopes).then(
      accessToken => {
        return accessToken
      },
      error => {
        return this.app
          .acquireTokenPopup(this.applicationConfig.graphScopes)
          .then(
            accessToken => {
              return accessToken
            },
            err => {
              console.error(err)
            }
          )
      }
    )
  }

  getGraphUserInfo(token) {
    const headers = new Headers({ Authorization: `Bearer ${token}` });
    const options = {
      headers
    };
    return fetch(`${this.graphUrl}`, options)
      .then(response => response.json())
      .catch(response => {
        throw new Error(response.text());
      });
  }

  // Utility
  getUser() {
    return this.app.getUser()
  }
}

export default new AuthenticationService()
