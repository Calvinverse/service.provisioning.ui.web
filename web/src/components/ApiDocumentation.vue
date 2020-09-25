<template>
  <div>
    <!-- For render-style=read you should try to provide height of rapi-doc element as it needs to calculate scroll positions-->
    <rapi-doc
      v-for="(doc, name) in obj"
      :key="name"
      :spec-url="doc.href"
      render-style="read"
      style="height:100vh; width:100%"
    ></rapi-doc>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import 'rapidoc'

export default Vue.extend({
  name: 'ApiDocumentation',

  data: () => ({
    obj: [
      {
        text: 'Provisioning',
        href: '/api/v1/doc'
      }
    ]
  })
})
</script>
