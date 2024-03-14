<template>
  <div class="flex items-start
              flex-col gap-4 p-4 
              rounded-lg overflow-y-auto 
              max-h-[800px] shadow-2xl 
            
              border border-gray-700
              font-inter">

    <div class="flex flex-col gap-4">
      <h1 class="text-lg font-bold">From: {{ mailData.from }}</h1>
      <div>
        <span class="font-bold">To: </span>
        <span class="text-gray-700">{{ mailData.to }}</span>
      </div>
      <div>
        <span class="font-bold">Subject: </span>
        <span class="text-gray-700">{{ mailData.subject }}</span>
      </div>
      <div>
        <span class="font-bold">Date: </span>
        <span class="text-gray-700">{{ mailData.date }}</span>
      </div>
      <div class="bg-gray-100 p-4 rounded-lg">

        <div>
          <span class="inline-block mx-1" v-for="(word, index) in separated_word" :key="index">
            <span v-if="searched_value && word.toLowerCase().includes(searched_value.toLowerCase())" class="bg-yellow-400">{{ word }}</span>
            <span v-else>{{ word }}</span>
          </span>
        </div>    

      </div>
    </div>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        separated_word:null,
      }
    },
    props: {
      mailData : {
        type: Object,
        required: true
      },
      searched_value:{
        type: String,
        required: false
      }
    },
    methods:{
    },
    mounted(){},
    watch: {
      mailData: {
        immediate: true, 
        handler(newVal) {
          if (newVal && newVal.content) {
            this.separated_word = newVal.content.split(" ");
          }
        }
      }
    },
    components:{}
  }
</script>

<style scoped>
@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-15px);
  }
}

.animate-bounce {
  animation: bounce 2s infinite;
}

.font-inter {
  font-family: 'Inter', sans-serif;
}
</style>
