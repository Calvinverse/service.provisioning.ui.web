<template>
  <v-container fill-height>
    <v-container
      >
      <v-row v-for="(item, index) in environments" :key="item.id" :data-index="index">
        <v-col>
          <v-card
            class="rounded-xl"
            color="primary"
            outlined
            @click="navigateToDetail(item.id)">
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
                <v-spacer />
                <v-col>
                  <h3>Resources</h3>
                  <span>{{ item.resources.length }}</span>
                </v-col>
              </v-row>
            </v-container>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
    <v-bottom-navigation>
      <v-btn>
        <span>Add</span>
        <v-icon>mdi-plus-thick</v-icon>
      </v-btn>

      <v-btn>
        <span>Delete</span>
        <v-icon>mdi-delete</v-icon>
      </v-btn>
    </v-bottom-navigation>
  </v-container>
</template>

<script lang="ts">

import Vue from 'vue'
import {
  namespace
} from 'vuex-class'
import Component from 'vue-class-component'
import { Channel, EventBus, Listener, Topic } from 'estacion/lib'
import { EventBusService } from '../services/EventBusService'
import { Constants } from '../types'
import { Environment } from '../services/EnvironmentService'

const environmentModule = namespace('environment')

@Component
export default class EnvironmentsList extends Vue {
  private bus: EventBus
  private authenticationChannel: Channel
  private loginTopic: Topic
  private logoutTopic: Topic

  constructor () {
    super()

    this.bus = EventBusService.getInstance().getBus()
    this.authenticationChannel = this.bus.channel(Constants.authenticationChannel)
    this.loginTopic = this.authenticationChannel.topic(Constants.userLoginTopic)
    this.logoutTopic = this.authenticationChannel.topic(Constants.userLogoutTopic)

    const loginListener: Listener = event => {
      this.getAll()
    }
    this.loginTopic.addListener(loginListener)

    const logoutListener: Listener = event => {
      this.clear()
    }
    this.logoutTopic.addListener(logoutListener)
  }

  @environmentModule.Action('clear') clear: any
  @environmentModule.Action('getAll') getAll: any
  @environmentModule.Action('delete') delete: any
  @environmentModule.Getter('environments') environments!: Environment[]
  @environmentModule.Getter('hasAny') hasAny!: boolean

  environmentStatusIcon (environment: Environment) {
    return environment.status() === 'ok' ? 'mdi-checkbox-marked-circle' : 'mdi-close-circle'
  }

  environmentStatusColor (environment: Environment) {
    return environment.status() === 'ok' ? 'success' : 'error'
  }

  navigateToDetail (environmentID: string): void {
    this.$router.push({
      name: 'Environment',
      params: {
        id: environmentID
      }
    })
  }
}
</script>
