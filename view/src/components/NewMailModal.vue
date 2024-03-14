<template>
  <div v-if="showModal" class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-50">
    <div class="bg-white p-6 rounded shadow-md relative">
      <button @click="close_modal" class="absolute top-0 right-0 m-4">&times;</button>
      <form @submit.prevent="submitForm">
        <input type="text" placeholder="From:" v-model="this.from" class="block mb-4">
        <input type="text" placeholder="To:" v-model="this.to" class="block mb-4">
        <input type="text" placeholder="Subject:" v-model="this.subject" class="block mb-4">
        <input type="date" placeholder="Date" v-model="this.date" class="block mb-4">
        <input type="text" placeholder="Content" v-model="this.content" class="block mb-4">
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
    },
    sendEmail: {
      type: Function,
      required: true
    },
    close_modal: {
      type: Function,
      required: true
    }
  },
  data() {
    return {
      to: "",
      from: [],
      subject: "",
      date: new Date().toISOString().split('T')[0],
      content: ""
    }
  },
  methods: {
    async submitForm() {
      const currentTime = new Date().toISOString().split('T')[1].slice(0, 8)
      const email = { 
        from: this.from, 
        to: this.to.split(","), 
        subject: this.subject, 
        date: this.date + 'T' + currentTime + "-08:00",
        content: this.content }
      await this.sendEmail(email)
    }
  }
}
</script>
<style>
  
</style>