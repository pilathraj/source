<template>
  <div class="q1">
    <section>
        <div class="container">
           <div class="columns is-centered">
               <div class="column is-half">
                <form @submit="formSubmit">
                    <b-field>
                        <b-input placeholder="Column to Start"
                            type="input"
                            icon="magnify"  v-model="start">
                        </b-input>
                        <b-input placeholder="No of rows"
                            type="number"
                            icon="magnify"  v-model="rows">
                        </b-input>
                        <b-input placeholder="No of columns"
                            type="number"
                            icon="magnify"  v-model="columns">
                        </b-input>
                        <p class="control">
                            <b-button type="is-primary" @click="formSubmit" label="Submit" />
                        </p>
                    </b-field> 
                </form> 
            </div>
           </div>
        </div>
    </section>
    <div>&nbsp;</div>
    <section>
        <div class="container">
           <div class="columns is-centered">
               <div class="column is-half">        
                    <div>{{errMsg}}</div>
                    <table class="table is-bordered">
                        <tbody>
                            <tr v-for= "(rows, rowIdx) in data" :key = "rowIdx">
                                <td v-for= "(col, colIdx) in rows" :key = "colIdx">
                                {{col}}
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </section>
  </div>    
</template>

<script>
export default {
  name: 'Q2',
  props: {
    msg: String
    
  },
  data() {
    return {
    start :"A",
    rows: 3,
    columns: 3,
    data: [],
    errMsg: ""
    };
  },
  methods: {
    formSubmit(e) {
        e.preventDefault();
        let currentObj = this;
        this.axios.post('http://localhost:5000/q2', {
            ColumnStart: this.start,
            Rows: parseInt(this.rows),
            Columns: parseInt(this.columns)
        })
        .then(function (response) { 
            if(response.data.Data) {  
            currentObj.data = response.data.Data;
            } else {
                currentObj.errMsg = response.data;
            }
        })
        .catch(function (error) {            
            currentObj.errMsg = error;
        });
    }
 } 
}
</script>