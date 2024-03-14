<template>
  <div class="flex p-4 relative">
    <NewMailModal :showModal="showModal" :sendEmail="sendEmail" :close_modal="close_modal"/>
    <div class="flex flex-col gap-4 w-3/4 pr-4">
      <div class="w-1/4">
        <button @click="open_modal" class="bg-blue-700 hover:bg-blue-500 text-white font-bold py-2 px-4 rounded">New</button>
      </div>
      <SearchData :search="search_value"/>
      <DataTable :data="this.mails" :query="this.query" :showMailIndex="updatedSelectedMail"/>
      <div class="mt-4">
        <Pagination :meta_data="this.pagination" :nextPage="nextPage" :prevPage="prevPage" :inputPage="inputPage"/>
      </div>
    </div>
    <div class="w-1/4">
      <MailData :mailData="this.selectedMail" :searched_value="finded_value"/>
    </div>
  </div>
</template>
<script>
import Pagination from "../components/Pagination.vue";
import DataTable from "../components/DataTable.vue"
import MailData from "../components/MailData.vue";
import SearchData from "../components/Search.vue";
import NewMailModal from "../components/NewMailModal.vue";
import { MailService } from "../services/email-service"

export default {
  created() {
    this.mailService = new MailService();
  },
  async mounted() {
    await this.search_data();
  },
  data() {
    return {
      query: "",
      size: 10,
      from: 0,
      sort: "-@timestamp",
      mails: [],
      pagination: {},
      selectedMail: {},
      finded_value: '',
      showModal: false,
      campo1: '',
      campo2: '',
      campo3: '',
      campo4: ''
    };
  },
  methods: {
    open_modal() {
      this.showModal = true
    },
    close_modal() {
      this.showModal = false
    },
    async search_data() {
      if(this.query !== "") {
        this.sort = "-date"
      }
      const request = {
        query: this.query,
        size: this.size,
        from: this.from,
        sort: this.sort
      };
      const { data, meta } = await this.mailService.search_data(request);
      this.mails = [...data]
      this.pagination = {...meta}
      this.selectedMail = this.mails[0]
    },
    updatedSelectedMail(index) {
      this.selectedMail = this.mails[index]
    },
    async prevPage() {
      if(!this.pagination.has_prev_page)
        return
      this.from -= this.size
      await this.search_data()
    },
    async nextPage() {
      if(!this.pagination.has_next_page)
        return
      this.from += this.size
      await this.search_data()
    },
    async search_value(searchValue) {
      const texto = this.query = searchValue;
      this.finded_value = texto;
        
      const request = {
        query: texto,
        size: this.size,
        from: 0,
        sort: "-date"
      };
    
      const { data, meta } = await this.mailService.search_data(request);
      this.mails = [...data]
      this.pagination = {...meta}
      this.selectedMail = this.mails[0]
    },
    async inputPage() {
      const { current_page, total_pages, items_per_page} = this.pagination
      if(current_page > total_pages)
      return
      this.from = (current_page - 1) * items_per_page
      await this.search_data()
    },
    async sendEmail(email) {
      console.log("Send", email)
      const response = await this.mailService.new_mail(email)
      console.log(response)
    }
  },
  components: { DataTable, MailData, Pagination, SearchData, NewMailModal }
}
</script>
<style lang="">
</style>