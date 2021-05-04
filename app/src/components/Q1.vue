<template>
  <div class="q1">
    <section>
       <form @submit="formSubmit">
          <b-field>
              <b-input placeholder="URL to Crawl..."
                  type="search"
                  icon="magnify" expanded v-model="url">
              </b-input>
              <p class="control">
                  <b-button type="is-primary" @click="formSubmit" label="Submit" />
              </p>
          </b-field> 
          </form>       
    </section>
    <div>&nbsp;</div>
    <section>
        <div class="has-text-centered">        
        <pre>{{errMsg}}</pre>
        <table class="table is-bordered">
          <tbody>
          <tr v-for= "(rows, rowIdx) in tableData" :key = "rowIdx">
            <td v-for= "(data, colIdx) in rows" :key = "colIdx">
              {{data}}
            </td>
          </tr>
        </tbody>
        </table>
        </div>
    </section>
  </div>    
</template>

<script>
export default {
  name: 'Q1',
  props: {
    msg: String
    
  },
  data() {
    return {
    url :"",
    data: {},
    errMsg: ""
    };
  },
  methods: {
    formSubmit(e) {
        e.preventDefault();
        let currentObj = this;
        this.axios.post('http://localhost:5000/q1', {
            Url: this.url            
        })
        .then(function (response) {            
            currentObj.data = response.data;
        })
        .catch(function (error) {             
            currentObj.errMsg = error;
        });
    }
 },
 "computed":{
   tableData(){
      let index = 0   
      const cols = 4
      const len = Object.keys(this.data).length
      const rows = Math.ceil (len/cols)
      // create a two-dimensional array of rows x columns
      const arrays = Array.from (new Array (rows), () =>new Array (cols) .fill (" "))
      
      Object.keys(this.data).forEach ((item, key) =>{
        const rowIdx = Math.floor (index/cols)
        const colIdx = index% cols
        arrays [rowIdx] [colIdx] = item+' - '+key
        index++
      })
      return arrays
    
   }
 },
}
</script>