<template>
  <div class="q1">
    <section>
        <div class="container">
           <div class="columns is-centered">
               <div class="column is-half"> 
                <form @submit="formSubmit">
                    <b-field>
                        <b-input placeholder="Date in YYYY-MM-DDTHH:NN"
                            type="input"
                            icon="magnify" v-model="date">
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
    <div style="margin:20px;"> </div>
    <section>
        <div class="has-text-centered">        
        {{Msg}}      
        </div>
    </section>
  </div>    
</template>

<script>
export default {
  name: 'Q4',
  props: {
    msg: String
    
  },
  data() {
    return {
    date :"2020-02-18T01:50",    
    Msg: ""
    };
  },
  methods: {
    formSubmit(e) {
        e.preventDefault();
        let currentObj = this;
        this.axios.post('http://localhost:5000/q4', {
            Date: this.date            
        })
        .then(function (response) {            
            currentObj.Msg = response.data;
        })
        .catch(function (error) {             
            currentObj.Msg = error;
        });
    }
 }, 
}
</script>