<template>
  <div class="flex flex-row gap-3">
    <input type="text" placeholder="Search..." v-model="searchValue"
          class="shadowed-text px-4 py-2 text-gray-700 bg-white border 
        border-gray-300 rounded-md focus:outline-none focus:border-blue-500">
     <button class="bg-teal-700 hover:bg-teal-500 text-white font-bold py-2 px-4 rounded " @click="search(this.searchValue)">Search</button>
     
  </div>
</template>


<script>
export default {
  props: {
    search: {
      type: Function,
      required: true
    }
  },
  data() {
    return {
      searchValue: '',
    };
  }
};
</script>


<style>
.shadowed-text { text-shadow: 0px 1px 0px black;}
</style>
