<template>
  <v-app
    id="main"
    :style="{background: $vuetify.theme.themes.dark.background}">
    <v-app-bar
      absolute
      app
      color="primary"
      prominent
    >
      <v-container
        fluid>
        <v-row
          align="vertical"
          >
          <v-col
            cols="1">
            <v-img
              alt="Vuetify Logo"
              class="shrink mr-2"
              contain
              :src="require('./assets/logo-white.png')"
              transition="scale-transition"
              width="80"
            />
          </v-col>
          <v-col>
            <h1>
              Environments
            </h1>
          </v-col>
          <v-spacer></v-spacer>
          <v-col
          cols="2">
            <UserLoginState />
        </v-col>
        </v-row>
      </v-container>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer
      color="primary darken-2"
      padless
    >
      <v-row
        justify="center"
        no-gutters
      >
        <v-btn
          v-for="link in data.footerItems"
          :key="link.title"
          color="white"
          rounded
          class="my-2"
          :to="link.href"
          text
        >
          {{ link.title }}
        </v-btn>
        <v-col
          class="primary darken-4 py-4 text-center white--text"
          cols="12"
        >
          <strong>Build on: </strong>{{ serviceInfo.buildtime }} — <strong>Version: </strong>{{ serviceInfo.version}}
        </v-col>
      </v-row>
    </v-footer>
  </v-app>

  <!--
    Footer:

    icon: categories by Marie Van den Broeck from the Noun Project

  -->
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import ServiceInformationService from './services/ServiceInformationService'

export class ServiceInfo {
  public buildtime: Date
  public revision: string
  public version: string

  constructor (
    bt: Date,
    r: string,
    v: string
  ) {
    this.buildtime = bt
    this.revision = r
    this.version = v
  }
}

@Component
export default class App extends Vue {
  private serviceInfo: ServiceInfo = new ServiceInfo(
    new Date('1900-01-01 01:01:01'),
    '123456789abcdef',
    '0.0.0'
  )

  private data: any = {
    footerItems: [
      { title: 'Home', href: '/' },
      { title: 'About Us', href: '/about' }
    ],
    menuItems: [
      { title: 'Dashboard', icon: 'mdi-view-dashboard' },
      { title: 'Photos', icon: 'mdi-image' },
      { title: 'About', icon: 'mdi-help-box' }
    ]
  }

  getServiceInfo () {
    ServiceInformationService.get()
      .then((response) => {
        this.serviceInfo = new ServiceInfo(
          new Date(response.data.buildtime),
          response.data.revision,
          response.data.version
        )
        console.log(response.data)
      })
      .catch(function (error) {
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          console.log(error.response.data)
          console.log(error.response.status)
          console.log(error.response.headers)
        } else if (error.request) {
          // The request was made but no response was received
          // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
          // http.ClientRequest in node.js
          console.log(error.request)
        } else {
          // Something happened in setting up the request that triggered an Error
          console.log('Error', error.message)
        }
        console.log(error.config)
      })
  }

  mounted () {
    console.log('Loading service info ...')
    this.getServiceInfo()
  }
}
</script>
