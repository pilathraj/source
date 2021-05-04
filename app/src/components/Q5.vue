<template>
  <div class="q5">
    <section>
        <div class="container">
           <div class="columns is-centered">
               <div class="column is-half"> 
                <form @submit="formSubmit">
                    <b-field>
                        <b-input placeholder="Input"
                            type="number"
                            icon="magnify" v-model="input">
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
  name: 'Q5',
  props: {
    msg: String
    
  },
  data() {
    return {
    input :"13",    
    Msg: ""
    };
  },
  methods: {
    formSubmit(e) {
        e.preventDefault();
        let currentObj = this;
        this.axios.get('http://localhost:5000/q5/'+this.input)
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