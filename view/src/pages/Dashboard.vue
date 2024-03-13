<template>
  <div class="flex p-4">
    <div class="flex flex-col gap-4 w-3/4 pr-4">
        <SearchData :search="search_value"/>
        <DataTable :data="this.mails" :query="this.query" :showMailIndex="updatedSelectedMail"/>
        <div class="mt-4">
          <Pagination :meta_data="this.pagination" :nextPage="nextPage" :prevPage="prevPage"/>
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
            mails: [],
            pagination: {},
            selectedMail: {},
            finded_value: ''
        };
    },
    methods: {
        async search_data() {
            const request = {
                query: this.query,
                size: this.size,
                from: this.from,
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
              from: this.from,
            };
        
            const { data, meta } = await this.mailService.search_data(request);
            this.mails = [...data]
            this.pagination = {...meta}
            this.selectedMail = this.mails[0]
        }

    },
    components: { DataTable, MailData, Pagination, SearchData }
}
</script>
<style lang="">
  
</style>