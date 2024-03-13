<template>
  <tr @click="showMail(index)">
    <td class="px-2 py-2 hover:bg-gray-100 border border-blue-950 font-semibold font-inter">{{ mail.from }}</td>
    <td class="px-2 py-2 hover:bg-gray-100 border border-blue-950 font-semibold font-inter">{{ mail.to }}</td>
    <td class="px-2 py-2 hover:bg-gray-100 border border-blue-950 font-semibold font-inter">{{ mail.subject }}</td>
  </tr>
</template>
<script>
export default {
  props: {
    mail: {
      type: Object,
      required: true
    },
    index: {
      type: Number,
      required: true
    },
    showMail: {
      type: Function,
      required: true
    }
  }
}
</script>
<style>

.font-inter {
  font-family: Arial, sans-serif;
}

</style>