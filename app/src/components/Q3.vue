<template> 
 <div class="container">
           <div class="columns is-centered">
               <div class="column is-half">
    <div>
        <table v-if="data.length!=0" class="table">
            <thead>
                <tr>
                    <th v-for="(column,index) in columns" :key="index">{{column.label}}</th>
                    <th>#</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(row,index_row) in data" :key="index_row" v-bind:class="{ editing: editId === row.id }" >
                    <td v-for="(column,index) in columns" :key="index"> {{row[column.field]}} </td>    
                    <td>
                        <p class="buttons">
                            <a class="button is-small is-primary" @click="editUser(row)">Edit</a>
                            <a class="button is-small is-danger" @click="deleteUser(row.id)">Delete</a>
                        </p>
                    </td>   
                </tr>
            </tbody>
        </table> 
        <div v-else class="notification is-danger">
            Warning !<strong> No Data to Show</strong>
        </div>
    </div>
    <div class="column">
        <a href="#" v-show="showAdd == false" @click="addUser()" class="button is-small is-danger" style="float:right;margin-top:-20px" >X</a>
        <UserInputForm :editId="editId" :user="user"></UserInputForm>
    </div>
</div>
           </div>
           </div>
</template>

<script>
import UserInputForm from './UserInputForm.vue';
    export default {
        components:{
            UserInputForm,
        },
        data() {
            return {
                showAdd:true,
                editId:'',
                user:{
                    loginId:'',
                    full_name:'',
                    enabled:false,
                },
                data: [
                    { 'id': "1", 'loginId': 'Jesse', 'fullname': 'Simmons', 'enabled': true },
                    { 'id': "2", 'loginId': 'John', 'fullname': 'Jacobs', 'enabled': false },
                    { 'id': "3", 'loginId': 'Tina', 'fullname': 'Gilbert', 'enabled': false },
                    { 'id': "4", 'loginId': 'Clarence', 'fullname': 'Flores', 'enabled': true },
                    { 'id': "5", 'loginId': 'Anne', 'fullname': 'Lee',  'enabled': true }
                ],
                columns: [
                    {
                        field: 'id',
                        label: 'ID'                       
                    },
                    {
                        field: 'loginId',
                        label: 'LoginId',                        
                    },
                    {
                        field: 'fullname',
                        label: 'Full Name'                       
                    },
                    {
                        field: 'enabled',
                        label: 'Enable'                                              
                    }                                    
                ]
            }
        },
       
        methods:{
            editUser(user){
                this.editId = user.id;
                this.showAdd = false;
                this.user = user;
            },
            addUser(){
                this.editId = "";
                this.showAdd = true;
                this.user = {
                    loginId:'',
                    full_name:'',
                    enabled:false}
            },
            deleteUser(user){
                console.log(user);
            }
        }
    }
</script>