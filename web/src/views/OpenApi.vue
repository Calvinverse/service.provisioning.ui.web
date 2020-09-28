<template>
  <div class="api-documentation">
    <ApiDocumentation />
  </div>
</template>

<script>
// @ is an alias to /src
import ApiDocumentation from '@/components/ApiDocumentation.vue'

export default {
  name: 'OpenApi',
  components: {
    ApiDocumentation
  }
}
</script>
