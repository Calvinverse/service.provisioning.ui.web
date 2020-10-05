import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default new Vuetify({
  icons: {
    iconfont: 'mdi'
  },

  theme: {
    themes: {
      dark: {
        primary: {
          base: '#276099',
          lighten1: '#3c6fa3',
          lighten2: '#527fad',
          lighten3: '#678fb7',
          lighten4: '#7d9fc1',
          lighten5: '#93afcc',
          darken1: '#235689',
          darken2: '#1f4c7a',
          darken3: '#1b436b',
          darken4: '#17395b',
          darken5: '#13304c'
        },
        secondary: '#7c7b6f',
        accent: '#2196f3',
        error: '#f44336',
        warning: '#ff9800',
        info: '#673ab7',
        success: '#4caf50',
        background: '#23231f'
      }
    },
    dark: true
  }
})
