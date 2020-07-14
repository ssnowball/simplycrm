<template>
  <div class="container new-width">
    <h3 v-if="!updateID">Create {{table}}</h3>
    <h3 v-if="updateID">Update {{table}}</h3>
    <form>
      <ul class="collapsible popout">
      <li>
        <div class="collapsible-header">Info - click to expand <div class="helper-text" v-if="$v.InfoValid.$invalid"> - *required fields</div></div>
        <div class="collapsible-body">
            <div class="row">
              <div class="input-field col s4">
                <input id="fIName" type="text" v-model.trim="$v.name.$model">
                <label for="fIName" v-bind:class="{ active: name }">Name</label>
                <div class="helper-text" v-if="$v.name.$invalid">*required</div>
              </div>
              <div class="input-field col s4">
                <input id="fICompaniesHouse" type="text" v-model.trim="$v.companiesHouse.$model">
                <label for="fICompaniesHouse" v-bind:class="{ active: companiesHouse }">Companies House</label>
                <div class="helper-text" v-if="$v.companiesHouse.$invalid">*required</div>
              </div>
              <div class="input-field col s4">
                <select id="fIType" v-model.trim="$v.cType.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropType" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value">
                {{ item.value }}
                </option>
                </select>
                <label>Type</label>
                <div class="helper-text" v-if="$v.cType.$invalid">*required</div>
              </div>
            </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header">Address - click to expand <div class="helper-text" v-if="$v.AddressValid.$invalid"> - *required fields</div></div>
        <div class="collapsible-body">
            <div class="row">
              <div class="input-field col s3">
                <input id="fALineOne" type="text" v-model.trim="$v.lineOne.$model">
                <label for="fALineOne" v-bind:class="{ active: lineOne }">Line One</label>
                <div class="helper-text" v-if="$v.lineOne.$invalid">*required</div>
              </div>
              
              <div class="input-field col s3">
                <input id="fALineTwo" type="text" v-model.trim="$v.lineTwo.$model">
                <label for="fALineTwo" v-bind:class="{ active: lineTwo }">Line Two</label>
                <div class="helper-text" v-if="$v.lineTwo.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fALineThree" type="text" v-model.trim="$v.lineThree.$model">
                <label for="fALineThree" v-bind:class="{ active: lineThree }">Line Three</label>
                <div class="helper-text" v-if="$v.lineThree.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fALineFour" type="text" v-model.trim="$v.lineFour.$model">
                <label for="fALineFour" v-bind:class="{ active: lineFour }">Line Four</label>
                <div class="helper-text" v-if="$v.lineFour.$invalid">*required</div>
              </div>
            </div>
            <div class="row">
              <div class="input-field col s3">
                <input id="fALineFive" type="text" v-model.trim="$v.lineFive.$model">
                <label for="fALineFive" v-bind:class="{ active: lineFive }">Line Five</label>
                <div class="helper-text" v-if="$v.lineFive.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fAPostcode" type="text" v-model.trim="$v.postcode.$model">
                <label for="fAPostcode" v-bind:class="{ active: postcode }">Postcode</label>
                <div class="helper-text" v-if="$v.postcode.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fACountry" type="text" v-model.trim="$v.country.$model">
                <label for="fACountry" v-bind:class="{ active: country }">Country</label>
                <div class="helper-text" v-if="$v.country.$invalid">*required</div>
              </div>
              
              <div class="input-field col s3">
                <input id="fAWThreeW" type="text" v-model.trim="$v.w3w.$model">
                <label for="fAWThreeW" v-bind:class="{ active: w3w }">What3Words</label>
                <div class="helper-text" v-if="$v.w3w.$invalid">*required</div>
              </div>
            </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header">Contact - click to expand <div class="helper-text" v-if="$v.ConactValid.$invalid"> - *required fields</div></div>
        <div class="collapsible-body">
          <div class="row">
            <div class="input-field col s6">
              <input id="fCName" type="text" v-model.trim="$v.contactName.$model">
              <label for="fCName" v-bind:class="{ active: contactName }">Name</label>
              <div class="helper-text" v-if="$v.contactName.$invalid">*required</div>
            </div>
            
            <div class="input-field col s6">
              <input id="fCPosition" type="text" v-model.trim="$v.contactPosition.$model">
              <label for="fCPosition" v-bind:class="{ active: contactPosition }">Position</label>
              <div class="helper-text" v-if="$v.contactPosition.$invalid">*required</div>
            </div>
          </div>
          
          <div class="row">            
            <div class="input-field col s6">
              <input id="fCNumber" type="text" v-model.trim="$v.contactNumber.$model">
              <label for="fCNumber" v-bind:class="{ active: contactNumber }">Number</label>
              <div class="helper-text" v-if="$v.contactNumber.$invalid">*required</div>
            </div>
            <div class="input-field col s6">
              <input id="fCEmail" type="text" v-model.trim="$v.contactEmail.$model">
              <label for="fCEmail" v-bind:class="{ active: contactEmail }">Email</label>
              <div class="helper-text" v-if="$v.contactEmail.$invalid">*required</div>
            </div>
          </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header" v-if="updateID">Extra - click to expand</div>
        <div class="collapsible-body">
          <div class="row">
            <div class="input-field col s4">
              <input type="text" id="fELeadCount" readonly v-model.trim="$v.leadCount.$model">
              <label for="fELeadCount" v-bind:class="{ active: leadCount }">Leads</label>
            </div>
            <div class="input-field col s2">
              <a class="waves-effect waves-light btn-small" id="filterLead" v-on:click="filterResults($event)" v-bind:class="{ disabled: leadCount > 0 ? false : true }">view</a>
            </div>

            <div class="input-field col s4">
              <input type="text" id="fEActivityCount" readonly v-model.trim="$v.activityCount.$model">
              <label for="fEActivityCount" v-bind:class="{ active: activityCount }">Activity</label>
            </div>
            <div class="input-field col s2">
              <a class="waves-effect waves-light btn-small" id="filterActivity" v-on:click="filterResults($event)" v-bind:class="{ disabled: activityCount > 0 ? false : true }">view</a>
            </div>
          </div>
        </div>
      </li>
    </ul>

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
import {required} from 'vuelidate/lib/validators';

export default {
  name: 'accountEntry',
  props: ["table", "updateID"],
  data() {
    return {
      dropType: [{ key: 'Small', value: 'Small' }, { key: 'Medium', value: 'Medium' }, { key: 'Large', value: 'Large' }],
      name: "",
      companiesHouse: "",
      cType: "",
      // parentID: "",
      // parentName: "",
      lineOne: "",
      lineTwo: "",
      lineThree: "",
      lineFour: "",
      lineFive: "",
      postcode: "",
      country: "",
      w3w: "",
      contactName: "",
      contactPosition: "",
      contactNumber: "",
      contactEmail: "",
      leadCount: "",
      activityCount: "",
      docID: "",
      createdDate: "",
      createdUsername: "",
    };
  },
  validations: {
    name: {
      required,
    },
    cType: {
      required,
    },
    lineOne: {
      required,
    },
    postcode: {
      required,
    },
    contactName: {
      required,
    },
    contactNumber: {
      required,
      // minValue: minValue(0.01)
    },
    companiesHouse: {
      // required,
    },
    // parentID: {
    //   // required,
    // },
    // parentName: {
    //   // required,
    // },
    lineTwo: {
      // required,
    },
    lineThree: {
      // required,
    },
    lineFour: {
      // required,
    },
    lineFive: {
      // required,
    },
    country: {
      // required,
    },
    w3w: {
      // required,
    },
    contactPosition: {
      // required,
    },
    contactEmail: {
      // required,
    },
    leadCount: {
      // required,
      // minValue: minValue(1),
    },
    activityCount: {
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
    GroupValid: ['name', 'cType', 'lineOne', 'postcode', 'contactName', 'contactNumber'],
    InfoValid: ['name', 'cType'],
    AddressValid: ['lineOne', 'postcode'],
    ConactValid: ['contactName', 'contactNumber'],
  },
  mounted: function () {
    M.AutoInit();
    
    if (this.updateID) {
      this.getUpdateData(this.updateID);
    }    
},
watch: {
  cType: function () {
    if (this.updateID) {
      var s = document.getElementById('fIType');
      var v = this.cType;
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
    filterResults(event) {
      this.$emit('filter', event.target.id.substring(6), 'list', this.name);
    },
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
            this.companiesHouse = result.Data.CompaniesHouse;
            this.cType = result.Data.CType;
            this.lineOne = result.Data.LineOne;
            this.lineTwo = result.Data.LineTwo;
            this.lineThree = result.Data.LineThree;
            this.lineFour = result.Data.LineFour;
            this.lineFive = result.Data.LineFive;
            this.postcode = result.Data.Postcode;
            this.country = result.Data.Country;
            this.w3w = result.Data.W3w;
            this.contactName = result.Data.ContactName;
            this.contactPosition = result.Data.ContactPosition;
            this.contactNumber = result.Data.ContactNumber;
            this.contactEmail = result.Data.ContactEmail;
            this.docID = result.Data.DocID;
            this.createdDate = result.Data.CreatedDate;
            this.createdUsername = result.Data.CreatedUsername;
            this.leadCount = result.Data.LeadCount;
            this.activityCount = result.Data.ActivityCount;

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
      args.CompaniesHouse = self.companiesHouse;
      args.CType = self.cType;
      args.LineOne = self.lineOne;
      args.LineTwo = self.lineTwo;
      args.LineThree = self.lineThree;
      args.LineFour = self.lineFour;
      args.LineFive = self.lineFive;
      args.Postcode = self.postcode;
      args.Country = self.country;
      args.W3w = self.w3w;
      args.ContactName = self.contactName;
      args.ContactPosition = self.contactPosition;
      args.ContactNumber = self.contactNumber;
      args.ContactEmail = self.contactEmail; 

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
      args.CompaniesHouse = self.companiesHouse;
      args.CType = self.cType;
      args.LineOne = self.lineOne;
      args.LineTwo = self.lineTwo;
      args.LineThree = self.lineThree;
      args.LineFour = self.lineFour;
      args.LineFive = self.lineFive;
      args.Postcode = self.postcode;
      args.Country = self.country;
      args.W3w = self.w3w;
      args.ContactName = self.contactName;
      args.ContactPosition = self.contactPosition;
      args.ContactNumber = self.contactNumber;
      args.ContactEmail = self.contactEmail;
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
