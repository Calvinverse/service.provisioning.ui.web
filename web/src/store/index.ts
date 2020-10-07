import Vue from 'vue'
import Vuex from 'vuex'
import { RootState } from './RootState'

import { environment } from './environment/index'
import { profile } from './profile/index'

Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    environment,
    profile
  }
})

export default store
