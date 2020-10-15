import Vue from 'vue'
import Vuex from 'vuex'

import { Environments } from './environment/index'
import Profile from './profile/index'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    environment: Environments,
    profile: Profile
  }
})

export default store
