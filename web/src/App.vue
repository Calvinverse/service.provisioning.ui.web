<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <v-img
          alt="Vuetify Logo"
          class="shrink mr-2"
          contain
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-logo-dark.png"
          transition="scale-transition"
          width="40"
        />

        <v-img
          alt="Vuetify Name"
          class="shrink mt-1 hidden-sm-and-down"
          contain
          min-width="100"
          src="https://cdn.vuetifyjs.com/images/logos/vuetify-name-dark.png"
          width="100"
        />
      </div>

      <v-spacer></v-spacer>

      <v-btn
        href="https://github.com/vuetifyjs/vuetify/releases/latest"
        target="_blank"
        text
      >
        <span class="mr-2">Latest Release</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
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

  components: {
  },

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
