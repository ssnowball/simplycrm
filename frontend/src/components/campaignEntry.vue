<template>
  <div class="container new-width">
    <h3 v-if="!updateID">Create {{table}}</h3>
    <h3 v-if="updateID">Update {{table}}</h3>
    <form>
      <div class="row">
        <div class="input-field col s6">
          <input id="fName" type="text" v-model.trim="$v.name.$model">
          <label for="fName" v-bind:class="{ active: name }">Name</label>
          <div class="helper-text" v-if="$v.name.$invalid">*required</div>
        </div>
        <div class="input-field col s6">
          <input id="fDescription" type="text" v-model.trim="$v.description.$model">
          <label for="fDescription" v-bind:class="{ active: description }">Description</label>
          <div class="helper-text" v-if="$v.description.$invalid">*required</div>
        </div>
</div>
<div class="row">
  <div class="input-field col s4">
    <select id="fStatus" v-model.trim="$v.status.$model">
      <option value="" disabled selected>Choose your option</option>
      <option v-for="item in dropStatus" 
              v-bind:key="item.key" 
              v-bind:value="item.value">
      {{ item.value }}
      </option>
    </select>
    <label>Status</label>
    <div class="helper-text" v-if="$v.status.$invalid">*required</div>
  </div>
  <div class="input-field col s4">
    <select id="fTypeChannel" v-model.trim="$v.typeChannel.$model">
      <option value="" disabled selected>Choose your option</option>
      <option v-for="item in dropTypeChannel" 
              v-bind:key="item.key" 
              v-bind:value="item.value">
      {{ item.value }}
      </option>
    </select>
    <label>Channel</label>
    <div class="helper-text" v-if="$v.typeChannel.$invalid">*required</div>
  </div>
  <div class="input-field col s4">
    <select id="fTypeMethod" v-model.trim="$v.typeMethod.$model">
      <option value="" disabled selected>Choose your option</option>
      <option v-for="item in dropTypeMethod" 
              v-bind:key="item.key" 
              v-bind:value="item.value">
      {{ item.value }}
      </option>
    </select>
    <label>Method</label>
    <div class="helper-text" v-if="$v.typeMethod.$invalid">*required</div>
  </div>
</div>
<div class="row">
  <div class="input-field col s4">
    <input id="fBudget" type="number" min=0  step="0.01" lang="nb" v-model.trim="$v.budget.$model">
    <label for="fBudget" v-bind:class="{ active: budget }">Budget</label>
    <div class="helper-text" v-if="$v.budget.$invalid">*required</div>
  </div>

  
  
  <div class="input-field col s4">
    <input id="fSDate" type="text" class="datepicker" v-model.lazy.trim="$v.startDate.$model">
    <label for="fSDate" v-bind:class="{ active: startDate }">Start Date</label>
    <div class="helper-text" v-if="$v.startDate.$invalid">*required</div>
  </div>
  <div class="input-field col s4">
    <input id="fEDate" type="text" class="datepicker" v-model.lazy.trim="$v.endDate.$model">
    <label for="fEDate" v-bind:class="{ active: endDate }">End Date</label>
    <div class="helper-text" v-if="$v.endDate.$invalid">*required</div>
  </div>

</div>

      <div class="row">
        <div class="input-field col s4">
          <input type="hidden" id="fDocID" disabled v-model.trim="$v.docID.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fcreatedDate" disabled v-model.trim="$v.createdDate.$model">
        </div>
        <div class="input-field col s4">
          <input type="hidden" id="fcreatedUsername" disabled v-model.trim="$v.createdUsername.$model">
        </div>
      </div>
        
        <a class="waves-effect waves-light btn entryButton" v-on:click="changePage">cancel</a>
        <a v-if="!updateID" class="waves-effect waves-light btn entryButton" id="btn-create" v-bind:class="{ disabled: $v.GroupValid.$invalid }" v-on:click="createData('btn-create')">create</a>
        <a v-if="updateID" data-target="modal-delete" class="btn modal-trigger delete-color entryButton" id="delete-trigger">delete</a>
        <a v-if="updateID" class="waves-effect waves-light btn entryButton" id="btn-update" v-bind:class="{ disabled: $v.GroupValid.$invalid }" v-on:click="updateData('btn-update')">update</a>

      </form>

      <div class="progress hide" id="loader-show">
        <div class="indeterminate"></div>
      </div>

      <!-- Modal Structure -->
      <div id="modal-delete" class="modal">
        <div class="modal-content">
          <h4>Confirm Delete</h4>
          <p>You are about to delete please confirm</p>
        </div>
        <div class="modal-footer">
          <a href="#!" class="modal-close waves-effect waves-green btn-flat">Cancel</a>
          <a class="waves-effect waves-light btn" id="btn-delete" v-on:click="deleteData()">Confirm</a>
        </div>
      </div>

  </div>
</template>

<script>
import M from 'materialize-css'
import {required, minValue} from 'vuelidate/lib/validators';

export default {
  name: 'campaignEntry',
  props: ["table", "updateID"],
  data() {
    return {
      dropStatus: [{ key: 'Running', value: 'Running' }, { key: 'Completed', value: 'Completed' }],
      dropTypeChannel: [{ key: 'Internal', value: 'Internal' }, { key: 'External', value: 'External' }],
      dropTypeMethod: [{ key: 'Gathering', value: 'Gathering' }, { key: 'Other', value: 'Other' }],
      name: "",
      description: "",
      status: "",
      typeChannel: "",
      typeMethod: "",
      budget: "",
      startDate: "",
      endDate: "",
      docID: "",
      createdDate: "",
      createdUsername: "",
    };
  },
  validations: {
    name: {
      required,
    },
    status: {
      required,
    },
    typeChannel: {
      required,
    },
    typeMethod: {
      required,
    },
    startDate: {
      required,
    },
    endDate: {
      required,
    },
    budget: {
      required,
      minValue: minValue(0.01)
    },
    description: {
      // required,
    },
    code: {
      // required,
    },
    docID: {
      // required,
    },
    createdDate: {
      // required,
    },
    createdUsername: {
      // required,
    },
    GroupValid: ['name', 'status', 'typeChannel', 'typeMethod', 'startDate', 'endDate', 'budget'],
  },
  mounted: function () {
    M.AutoInit();

    var elems = document.querySelectorAll('.datepicker');
    var options = {format:'dd/mm/yyyy'};
    M.Datepicker.init(elems, options);
    
    if (this.updateID) {
      this.getUpdateData(this.updateID);
    }    
},
watch: {
  status: function () {
      if (this.updateID) {
      var s = document.getElementById('fStatus');
      var v = this.status;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    typeChannel: function () {
      if (this.updateID) {
      var s = document.getElementById('fTypeChannel');
      var v = this.typeChannel;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
    }
    },
    typeMethod: function () {
      if (this.updateID) {
      var s = document.getElementById('fTypeMethod');
      var v = this.typeMethod;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    }, 
},
  methods: {
    changePage () {
      this.$emit('clicked', 'list')
    },
    deleteData: function() {
      document.getElementById("loader-show").classList.toggle('hide');

      var self = this;

      window.backend.deleteData(self.table, self.updateID).then(result => {

        if (result.Success === true) {
          M.toast({html: "Deleted Item", classes: 'rounded'});   
          this.changePage(); 
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }

        document.getElementById("loader-show").classList.toggle('hide');
          
        });

    },
    getUpdateData(docID) {

      window.backend.getUpdateData(this.table, docID).then(result => {       

        if (result.Data === null) {
            M.toast({html: result.ReturnErr, classes: 'rounded'});
          } else {

            this.name = result.Data.Name;
            this.description = result.Data.Description;
            this.status = result.Data.Status;
            this.typeChannel = result.Data.TypeChannel;
            this.typeMethod = result.Data.TypeMethod;
            this.budget = result.Data.Budget;
            this.startDate = result.Data.StartDate;
            this.endDate = result.Data.EndDate;
            this.docID = result.Data.DocID;
            this.createdDate = result.Data.CreatedDate;
            this.createdUsername = result.Data.CreatedUsername;

            M.updateTextFields();
          }
      });
    },
    createData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Description = self.description
      args.Status = self.status;
      args.TypeChannel = self.typeChannel;
      args.TypeMethod = self.typeMethod;
      args.Budget = self.budget; 
      args.StartDate = self.startDate; 
      args.EndDate = self.endDate; 

      // console.log(args);

      window.backend.createData(self.table, args).then(result => {
        if (result.Success === true) {
          M.toast({html: "Added Item", classes: 'rounded'});
          this.changePage();
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
        document.getElementById("loader-show").classList.toggle('hide');

      });

    },
    updateData(btnID) {
      var self = this;
      document.getElementById(btnID).classList.toggle('disabled');
      document.getElementById("loader-show").classList.toggle('hide');

      var args = {};
      args.Name = self.name;
      args.Description = self.description
      args.Status = self.status;
      args.TypeChannel = self.typeChannel;
      args.TypeMethod = self.typeMethod;
      args.Budget = self.budget; 
      args.StartDate = self.startDate; 
      args.EndDate = self.endDate;
      args.DocID = self.docID;
      args.CreatedDate = self.createdDate;
      args.CreatedUsername = self.createdUsername;

      window.backend.updateData(self.table, args).then(result => {
        if (result.Success === true) {
          M.toast({html: "Updated Item", classes: 'rounded'}); 
          this.changePage(); 
        } else {
          M.toast({html: result.ReturnErr, classes: 'rounded'});
        }
        document.getElementById("loader-show").classList.toggle('hide');

      });

    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
