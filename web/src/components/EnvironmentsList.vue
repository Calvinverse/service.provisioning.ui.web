<template>
  <v-container fill-height>
    <v-container
      >
      <v-row v-for="(item, index) in environments" :key="item.id" :data-index="index">
        <v-col>
          <v-card
            color="secondary"
            outlined>
            <v-container>
              <!-- Status image -->
              <v-row>
                <v-col>
                  <v-avatar>
                    <v-icon
                      :color="environmentStatusColor(item.status)"
                      x-large>
                      {{ environmentStatusIcon(item.status) }}
                    </v-icon>
                  </v-avatar>
                </v-col>
                <v-col>
                  <h2>{{ item.name }}</h2>
                  <span>{{ item.description }}</span>
                  <span>{{ item.version }}</span>
                </v-col>
                <v-spacer>
                </v-spacer>
                <v-col>
                  <h3>Resources</h3>
                  <span>{{ item.resources.length }}</span>
                </v-col>
              </v-row>
              <v-row>
                <v-col>
                  Created on:
                </v-col>
                <v-col>
                  {{ item.createdOn }}
                </v-col>
                <v-spacer />
              </v-row>
              <v-row>
                <v-col>
                  Destroy by:
                </v-col>
                <v-col>
                  {{ item.destroyBy }}
                </v-col>
                <v-spacer />
              </v-row>
              <v-row>
                <v-col>
                  <v-chip
                    v-for="(tag, tagIndex) in item.tags" :key="tag.name" :data-index="tagIndex"
                    color="success"
                    outlined
                  >
                    {{ tag.name }}: {{ tag.value }}
                  </v-chip>
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
import {
  Environment
} from '../store/environment'
import { Channel, EventBus, Listener, Topic } from 'estacion/lib'
import { EventBusService } from '../services/EventBusService'
import { Constants } from '../types'

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
      // clear
    }
    this.logoutTopic.addListener(logoutListener)
  }

  @environmentModule.Action('getAll') getAll: any
  @environmentModule.Action('delete') delete: any
  @environmentModule.Getter('environments') environments!: Environment[]
  @environmentModule.Getter('hasAny') hasAny!: boolean

  environmentStatusIcon (status: string) {
    return status === 'ok' ? 'mdi-checkbox-marked-circle' : 'mdi-close-circle'
  }

  environmentStatusColor (status: string) {
    return status === 'ok' ? 'success' : 'error'
  }
}
</script>
