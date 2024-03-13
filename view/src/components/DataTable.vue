<template>
  <table class="text-sm text-left shadow-2xl">
    <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400 border border-blue-950">
      <tr>
        <th class="px-2 py-2 hover:bg-gray-300">From</th>
        <th class="px-2 py-2 hover:bg-gray-300">To</th>
        <th class="px-2 py-2 hover:bg-gray-300">Subject</th>
      </tr>
    </thead>
    <tbody>
      <MailRow v-for="(mail, i) in data" :key="i" class="border-t align-top" :showMail="showMailIndex" :index="i" :mail="mail"/>
    </tbody>
  </table>
</template>
<script>
import MailRow from './MailRow.vue';

export default {
    props: {
        data: {
            type: Array,
            required: true
        },
        query: {
            type: String,
            required: false
        },
        showMailIndex: {
            type: Function,
            required: true
        }
    },
    methods: {
        formatContent(content) {
            // return content
            // if(this.query == "") 
            //   return content
            const queryPosition = content.toLowerCase().indexOf(this.query.toLowerCase());
            const paragraphs = content.split('\n\n');
            const formattedContent = paragraphs.map(paragraph => `${paragraph}`).join('');
            // return formattedContent
            const shortContent = (queryPosition > 40 ? "..." : "") + formattedContent.slice(queryPosition < 40 ? 0 : queryPosition - 40, queryPosition + 85) + "...";
            const highlightedContent = shortContent.replace(this.query, `<span style="background-color: yellow; font-weight: bold;">${this.query}</span>`);
            return highlightedContent;
        }
    },
    components: { MailRow }
}
</script>
<style lang="">

</style>