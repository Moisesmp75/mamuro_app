<template>
  <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
    <div class="bg-white p-6 rounded shadow-md relative">
      <button @click="showModal = false" class="absolute top-0 right-0 m-4">&times;</button>
      <form @submit.prevent="submitForm">
        <input type="text" placeholder="To:" v-model="to" class="block mb-4">
        <input type="text" placeholder="From:" v-model="from" class="block mb-4">
        <input type="text" placeholder="Subject:" v-model="subject" class="block mb-4">
        <input type="text" placeholder="Date" v-model="date" class="block mb-4">
        <input type="text" placeholder="Content" v-model="content" class="block mb-4">
        <button type="submit" class="py-2 px-4 bg-green-500 text-white rounded">Submit</button>
      </form>
    </div>
  </div>

</template>
<script>
export default {
  props: {
    showModal: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      to: "",
      from: [],
      subject: "",
      date: "",
      content: ""
    }
  }
}
</script>
<style>
  
</style>