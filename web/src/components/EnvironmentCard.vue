<template>
  <!--
    - Environment description at the top
    - Templates that were used
    - List of resource groups with their status
    - Way to update resources per group
    - Graph of resources?
  -->
  <v-container>
    <!--
      Environment info
    -->
    <v-row>
      <v-col>
        <v-card
          class="rounded-t-xl"
          color="primary">
          <v-container>
            <v-row>
              <v-col
                align-self="center"
                cols="1">
                <v-avatar
                  align="center"
                  justify="center">
                  <v-icon
                    :color="environmentStatusColor(item)"
                    x-large>
                    {{ environmentStatusIcon(item) }}
                  </v-icon>
                </v-avatar>
              </v-col>
              <v-col>
                <v-container
                  class="pa-0">
                  <v-row
                    class="pa-0 pb-2">
                    <v-col
                      class="pa-0">
                      <h2>{{ item.name }}</h2>
                      <span>{{ item.description }}</span>
                    </v-col>
                  </v-row>
                  <v-divider></v-divider>
                  <v-row
                    class="pa-0 pt-2">
                    <v-col
                      class="pa-0"
                      cols="3">
                      Version:
                    </v-col>
                    <v-col
                      class="pa-0">
                      {{ item.version }}
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col
                      class="pa-0"
                      cols="3">
                      Created on:
                    </v-col>
                    <v-col
                      class="pa-0">
                      {{ item.createdOn.toISOString() }}
                    </v-col>
                  </v-row>
                  <v-row>
                    <v-col
                      class="pa-0"
                      cols="3">
                      Destroy by:
                    </v-col>
                    <v-col
                      class="pa-0">
                      {{ item.destroyBy.toISOString() }}
                    </v-col>
                  </v-row>
                  <v-row>
                  <v-col
                    class="pa-0 pt-2">
                    <v-chip
                      class="mr-3 mb-3"
                      v-for="(tag, tagIndex) in item.tags" :key="tag.name" :data-index="tagIndex"
                      color="info"
                    >
                      {{ tag.name }}: {{ tag.value }}
                    </v-chip>
                  </v-col>
                </v-row>
                </v-container>
              </v-col>
            </v-row>
          </v-container>
        </v-card>
      </v-col>
    </v-row>

    <!-- Resources -->
    <v-row>
      <v-col>
        <v-data-table
          color="primary">

        </v-data-table>
      </v-col>
    </v-row>

    <!-- Templates -->

    <!-- Graph -->
    <v-row>

    </v-row>
  </v-container>
</template>

<script lang="ts">
import Vue from 'vue'
import {
  namespace,
  State
} from 'vuex-class'
import Component from 'vue-class-component'
import { Environment } from '../services/EnvironmentService'
import { Channel, EventBus, Listener, Topic } from 'estacion/lib'
import { EventBusService } from '../services/EventBusService'
import { Constants, TagDefinition } from '../types'

const environmentModule = namespace('environment')

@Component
export default class EnvironmentCard extends Vue {
  private bus: EventBus
  private authenticationChannel: Channel
  private loginTopic: Topic
  private logoutTopic: Topic

  item: Environment

  constructor () {
    super()

    this.bus = EventBusService.getInstance().getBus()
    this.authenticationChannel = this.bus.channel(Constants.authenticationChannel)
    this.loginTopic = this.authenticationChannel.topic(Constants.userLoginTopic)
    this.logoutTopic = this.authenticationChannel.topic(Constants.userLogoutTopic)

    const loginListener: Listener = event => {
      this.get()
    }
    this.loginTopic.addListener(loginListener)

    const logoutListener: Listener = event => {
      // this.clear()
    }
    this.logoutTopic.addListener(logoutListener)

    this.item = new Environment(
      '',
      '',
      '',
      new Date(),
      new Date(),
      new Array<string>(),
      '',
      new Array<TagDefinition>(),
      '')
  }

  @environmentModule.Action('get') get: any
  @environmentModule.Action('delete') delete: any
  @environmentModule.Getter('environment') environment!: (environmentID: string) => Environment

  environmentStatusIcon (environment: Environment) {
    return environment.status() === 'ok' ? 'mdi-checkbox-marked-circle' : 'mdi-close-circle'
  }

  environmentStatusColor (environment: Environment) {
    return environment.status() === 'ok' ? 'success' : 'error'
  }

  created () {
    this.loadEnvironment()
  }

  loadEnvironment () {
    this.item = this.environment(this.$route.params.id)
  }
}
</script>
