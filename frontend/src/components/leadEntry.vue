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
              <div class="input-field col s3">
                <input id="fIName" type="text" v-model.trim="$v.name.$model">
                <label for="fIName" v-bind:class="{ active: name }">Name</label>
                <div class="helper-text" v-if="$v.name.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fIDescription" type="text" v-model.trim="$v.description.$model">
                <label for="fIDescription" v-bind:class="{ active: description }">Description</label>
                <div class="helper-text" v-if="$v.description.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fIType" v-model.trim="$v.sType.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropType" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value">
                {{ item.value }}
                </option>
                </select>
                <label>Type</label>
                <div class="helper-text" v-if="$v.sType.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fISubType" v-model.trim="$v.sSubType.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropSubType" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == sSubType">
                {{ item.value }}
                </option>
                </select>
                <label>Type</label>
                <div class="helper-text" v-if="$v.sSubType.$invalid">*required</div>
              </div>
            </div>
            <div class="row">
              <div class="input-field col s11">
                    <input id="fIAccountName" type="text" v-model.trim="$v.accountName.$model" readonly>
                    <input id="fIAccountID" type="hidden" disabled v-model.trim="$v.accountID.$model">
                    <label for="fIAccountName" v-bind:class="{ active: accountName }">Account Name</label>
                    <div class="helper-text" v-if="$v.accountName.$invalid">*required</div>
              </div>
              <div class="input-field col s1">
                    <a class="waves-effect waves-light btn-small" id="account-find" v-on:click="getAccount('Account')">find</a>
              </div>
            </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header">Detail - click to expand <div class="helper-text" v-if="$v.DetailValid.$invalid"> - *required fields</div></div>
        <div class="collapsible-body">
            <div class="row">
              <div class="input-field col s3">
                <select id="fDStatus" v-model.trim="$v.status.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropStatus" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == status">
                {{ item.value }}
                </option>
                </select>
                <label>Status</label>
                <div class="helper-text" v-if="$v.status.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fDStage" v-model.trim="$v.stage.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropStage" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == stage">
                {{ item.value }}
                </option>
                </select>
                <label>Stage</label>
                <div class="helper-text" v-if="$v.stage.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fDLeadsourceType" v-model.trim="$v.leadsourceType.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropLeadsourceType" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == leadsourceType">
                {{ item.value }}
                </option>
                </select>
                <label>Leadsource Type</label>
                <div class="helper-text" v-if="$v.leadsourceType.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fDLeadsourceChannel" v-model.trim="$v.leadsourceChannel.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropLeadsourceChannel" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == leadsourceChannel">
                {{ item.value }}
                </option>
                </select>
                <label>Leadsource Channel</label>
                <div class="helper-text" v-if="$v.leadsourceChannel.$invalid">*required</div>
              </div>
            </div>
            <div class="row">
              <div class="input-field col s3">
                <input id="fDSaleValue" type="number" v-model.trim="$v.saleValue.$model" readonly class="tooltipped" data-position="top" data-tooltip="Add value by adding products below">
                <label for="fDSaleValue" v-bind:class="{ active: saleValue }">Sale Value</label>
                <div class="helper-text" v-if="$v.saleValue.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fAEstCloseDate" type="text" class="datepicker" v-model.lazy.trim="$v.estCloseDate.$model">
                <label for="fAEstCloseDate" v-bind:class="{ active: estCloseDate }">Est Close Date</label>
                <div class="helper-text" v-if="$v.estCloseDate.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <input id="fAActualCloseDate" type="text" class="datepicker" v-model.lazy.trim="$v.actualCloseDate.$model">
                <label for="fAActualCloseDate" v-bind:class="{ active: actualCloseDate }">Actual Close Date</label>
                <div class="helper-text" v-if="$v.actualCloseDate.$invalid">*required</div>
              </div>
              <div class="input-field col s3">
                <select id="fACloseReason" v-model.trim="$v.closeReason.$model">
                  <option value="" disabled selected>Choose your option</option>
                  <option v-for="item in dropCloseReason" 
                        v-bind:key="item.key" 
                        v-bind:value="item.value"
                        v-bind:selected="item.value == closeReason">
                {{ item.value }}
                </option>
                </select>
                <label>Close Reason</label>
                <div class="helper-text" v-if="$v.closeReason.$invalid">*required</div>
              </div>
            </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header">Products - {{pCounter}} - click to expand </div>
        <div class="collapsible-body" >
          <div class="row">
            <a class="waves-effect waves-light btn" id="product-add" v-on:click="getProducts()">add</a>
          </div>
          <div :key="componentKey">
          <div class="row" v-for="(item, index) in products" v-bind:key="item.UUID">
            <input type="hidden" id="fPUUID" readonly :value="item.UUID">
            <input type="hidden" id="fPDocID" readonly :value="item.DocID">


            <div class="input-field col s3">
              <input type="text" id="fPName" :value="item.Name" readonly>
              <label for="fPName" v-bind:class="{ active: item.Name }">Name</label>
            </div>
            <div class="input-field col s3">
              <input type="text" id="fPCategory" :value="item.Category" readonly>
              <label for="fPCategory" v-bind:class="{ active: item.Category }">Category</label>
            </div>
            <div class="input-field col s1">
              <input type="number" id="fPCost" :value="item.Cost" readonly>
              <label for="fPCost" v-bind:class="{ active: item.Cost }">Cost</label>
            </div>
            <div class="input-field col s1">
              <input type="number" id="fPQuantity" :value="item.Quantity" v-on:change="updateTotal(index, $event)">
              <label for="fPQuantity" v-bind:class="{ active: item.Quantity }">Quantity</label>
            </div>
            <div class="input-field col s1">
              <input type="number" id="fPDiscount" :value="item.Discount" v-on:change="updateTotal(index, $event)">
              <label for="fPDiscount" v-bind:class="{ active: item.Discount }">Discount</label>
            </div>
            <div class="input-field col s1">
              <input type="number" id="fPSale" :value="item.Sale" readonly>
              <label for="fPSale" v-bind:class="{ active: item.Sale }">Sale</label>
            </div>

            <div class="input-field col s1">
              <a class="waves-effect waves-light btn" v-on:click="removeItem(item.Table, index)">remove</a>
            </div>
          </div>
          </div>
        </div>
      </li>
      <li>
        <div class="collapsible-header">Campaigns - {{cCounter}} - click to expand </div>
        <div class="collapsible-body">
          <div class="row">
            <a class="waves-effect waves-light btn" id="campaign-add" v-on:click="getCampaigns()">add</a>
          </div>
          <div class="row" v-for="(item, index) in campaigns" v-bind:key="item.key">

            <input type="hidden" id="fCUUID" readonly :value="item.UUID">
            <input type="hidden" id="fCDocID" readonly :value="item.DocID">

            <div class="input-field col s6">
              <input type="text" id="fCName" :value="item.Name">
              <label for="fCName" v-bind:class="{ active: item.Name }">Campaign Name</label>
            </div>

            <div class="input-field col s4">
              <p>
                <label>
                  <input type="checkbox"  v-bind:class="{ checked: item.Successful }"/>
                  <span>Campaign Successful</span>
                </label>
              </p>
            </div>

            <div class="input-field col s2">
              <a class="waves-effect waves-light btn" v-on:click="removeItem(item.Table, index)">remove</a>
            </div>

          </div>
        </div>
      </li>
      
      <li>
        <div class="collapsible-header">Extra - click to expand</div>
        <div class="collapsible-body">
          <div class="row">
            <div class="input-field col s2">
              <input type="text" id="fEActivityCount" readonly v-model.trim="$v.activityCount.$model">
              <label for="fEActivityCount" v-bind:class="{ active: activityCount }">Activity</label>
            </div>
            <div class="input-field col s2">
              <a class="waves-effect waves-light btn-small" id="filterActivity" v-on:click="filterResults($event)" v-bind:class="{ disabled: activityCount > 0 ? false : true }">view</a>
            </div>


            <div class="input-field col s5">
              <input type="text" id="fEQuoteID" v-model.trim="$v.quoteID.$model">
              <label for="fEQuoteID" v-bind:class="{ active: quoteID }">Quote ID</label>
            </div>
            <div class="input-field col s5">
              <input type="text" id="fEInvoiceID" v-model.trim="$v.invoiceID.$model">
              <label for="fEInvoiceID" v-bind:class="{ active: invoiceID }">Invoice ID</label>
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
      <div id="modal-account" class="modal bottom-sheet">
        <div class="modal-content">
          <a data-target="modal-account" class="btn modal-close">close</a> 
          <a class="waves-effect waves-light btn-small" v-on:click="removeAccount()" v-bind:class="{ disabled: !accountName }">Clear Selected</a> 
          <a class="waves-effect waves-light btn-small" v-on:click="filterResultsAccount('Account')" v-bind:class="{ disabled: !accountName }" >View Selected</a>
          <div style="margin-bottom: 10px">
          <el-row>
            <el-col :span="6">
              <el-input v-model="filters[0].value" placeholder="Search"></el-input>
            </el-col>
          </el-row>
        </div>
          <data-tables :data="accountData" :pagination-props="{ pageSizes: [5, 10, 15] }" :filters="filters" :action-col="actionCol">
            <el-table-column v-for="title in titles" :prop="title.prop" :label="title.label" :key="title.label" sortable="custom">
            </el-table-column>
          </data-tables>

        </div>
      </div>

      <!-- Modal Structure -->
      <div id="modal-product" class="modal bottom-sheet">
        <div class="modal-content">
          <button data-target="modal-product" class="btn modal-close">close</button>
          <div style="margin-bottom: 10px">
          <el-row>
            <el-col :span="6">
              <el-input v-model="filters[0].value" placeholder="Search"></el-input>
            </el-col>
          </el-row>
        </div>
          <data-tables :data="productData" :pagination-props="{ pageSizes: [5, 10, 15] }" :filters="filters" :action-col="actionCol">
            <el-table-column v-for="title in titles" :prop="title.prop" :label="title.label" :key="title.label" sortable="custom">
            </el-table-column>
          </data-tables>
        </div>
      </div>

      <!-- Modal Structure -->
      <div id="modal-campaign" class="modal bottom-sheet">
        <div class="modal-content">
          <button data-target="modal-campaign" class="btn modal-close">close</button>
          <div style="margin-bottom: 10px">
          <el-row>
            <el-col :span="6">
              <el-input v-model="filters[0].value" placeholder="Search"></el-input>
            </el-col>
          </el-row>
        </div>
          <data-tables :data="campaignData" :pagination-props="{ pageSizes: [5, 10, 15] }" :filters="filters" :action-col="actionCol">
            <el-table-column v-for="title in titles" :prop="title.prop" :label="title.label" :key="title.label" sortable="custom">
            </el-table-column>
          </data-tables>
        </div>
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
  name: 'leadEntry',
  props: ["table", "updateID"],
  data() {
    return {
      dropType: [{key: 'Sales', value: 'Sales' }, 
                { key: 'Marketing', value: 'Marketing' }],
      dropSubType: [{ key: 'New', value: 'New' }, 
                    { key: 'AddOn', value: 'AddOn' }],
      dropStatus: [{ key: 'Open', value: 'Open' }, 
                    { key: 'Won', value: 'Won' }, 
                    { key: 'Lost', value: 'Lost' }],
      dropStage: [{ key: '1-Identified', value: '1-Identified' }, 
                  { key: '2-Qualified', value: '2-Qualified' }, 
                  { key: '3-Quoted', value: '3-Quoted' }, 
                  { key: '4-Negotiation', value: '4-Negotiation' }, 
                  { key: '5-Sold', value: '5-Sold' }],
      dropLeadsourceType: [{ key: 'Inbound', value: 'Inbound' }, 
                            { key: 'Outbound', value: 'Outbound' }],
      dropLeadsourceChannel: [{ key: 'Email', value: 'Email' }, 
                              { key: 'Phone', value: 'Phone' }, 
                              { key: 'LiveChat', value: 'LiveChat' }, 
                              { key: 'Transfer', value: 'Transfer' }],
      dropCloseReason: [{ key: 'Price', value: 'Price' }, 
                        { key: 'Competitor', value: 'Competitor' }, 
                        { key: 'Length', value: 'Length' }, 
                        { key: 'Sales', value: 'Sales' }],
      products: [],
      productData: [],
      titles: [],
      filters: [
        {
          prop: '',
          value: ''
        }
      ],
      actionCol: {
        label: 'Actions',
        props: {
          align: 'center',
        },
        buttons: [{
          props: {
            type: 'primary',
          },
          handler: row => {
            this.selectedSubItem(row);
          },
          label: 'Select'
        }]
      },
      campaigns: [],
      campaignData: [],
      name: "",
      description: "",
      sType: "",
      sSubType: "",
      accountData: [],
      accountSearch: "",
      accountID: "",
      accountName: "",
      status: "",
      stage: "",
      leadsourceType: "",
      leadsourceChannel: "",
      saleValue: "",
      estCloseDate: "",
      actualCloseDate: "",
      closeReason: "",
      pCounter: 0,
      cCounter: 0,
      quoteID: "",
      invoiceID: "",
      activityCount: "",
      docID: "",
      createdDate: "",
      createdUsername: "",
      componentKey: 0,
    };
  },
  validations: {
    name: {
      required,
    },
    sType: {
      required,
    },
    sSubType: {
      required,
    },
    accountID: {
      required,
    },
    accountName: {
      required,
    },
    status: {
      required,
    },
    stage: {
      required,
    },
    leadsourceType: {
      required,
    },
    leadsourceChannel: {
      required,
    },
    saleValue: {
      required,
    },
    estCloseDate: {
      required,
    },
    pCounter: {
      // required,
    },
    cCounter: {
      // required,
    },
    description: {
      // required,
    },
    // resellerID: {
    //   // required,
    // },
    // resellerName: {
    //   // required,
    // },
    actualCloseDate: {
      // required,
    },
    closeReason: {
      // required,
    },
    quoteID: {
      // required,
    },
    invoiceID: {
      // required,
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
    GroupValid: ['name', 'sType', 'sSubType', 'accountID', 'accountName', 'status', 'stage', 'leadsourceType', 'leadsourceChannel', 'saleValue', 'estCloseDate'],
    InfoValid: ['name', 'sType', 'sSubType', 'accountID', 'accountName'],
    DetailValid: ['status', 'stage', 'leadsourceType', 'leadsourceChannel', 'saleValue', 'estCloseDate'],
    // ProductValid: ['name', 'sType', 'sSubType', 'accountID', 'accountName', 'status', 'stage', 'leadsourceType', 'leadsourceChannel', 'saleValue', 'estCloseDate'],
    // CampaignValid: ['name', 'sType', 'sSubType', 'accountID', 'accountName', 'status', 'stage', 'leadsourceType', 'leadsourceChannel', 'saleValue', 'estCloseDate'],
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
    products: function () {
      this.pCounter = this.products.length;
    },
    campaigns: function () {
      this.cCounter = this.campaigns.length;
    },
    sType: function () {
      if (this.updateID) {
      var s = document.getElementById('fIType');
      var v = this.sType;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    sSubType: function() {
      if (this.updateID) {
      var s = document.getElementById('fISubType');
      var v = this.sSubType;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    status: function() {
      if (this.updateID) {
    var s = document.getElementById('fDStatus');
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
    stage: function() {
      if (this.updateID) {
    var s = document.getElementById('fDStage');
      var v = this.stage;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    leadsourceType: function() {
      if (this.updateID) {
var s = document.getElementById('fDLeadsourceType');
      var v = this.leadsourceType;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    leadsourceChannel: function() {
      if (this.updateID) {
var s = document.getElementById('fDLeadsourceChannel');
      var v = this.leadsourceChannel;
      for ( var i = 0; i < s.options.length; i++ ) {
        if ( s.options[i].text == v ) {
            s.options[i].selected = true;
            M.FormSelect.init(s);
            return;
        }
    }
      }
    },
    closeReason: function() {
      if (this.updateID) {
var s = document.getElementById('fACloseReason');
      var v = this.closeReason;
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
    filterResultsAccount(event) {
      this.$emit('filter', event, 'list', this.accountName);
    },
    updateTotal(index, event) {
      var targ = event.target.id.substring(2);

      if (targ == "Quantity") {
        this.products[index].Quantity = parseInt(event.target.value);
      } else {
        this.products[index].Discount = parseInt(event.target.value);
      }

      var curPodSale = this.products[index].Sale;

      this.products[index].Sale = (this.products[index].Cost * (this.products[index].Quantity == null ? 0 : this.products[index].Quantity) ) - (this.products[index].Discount == null ? 0 : this.products[index].Discount) ;
    
      this.componentKey += 1;  

      var old = parseInt(this.saleValue == null ? 0 : this.saleValue);
      var add = (this.products[index].Sale == null ? 0 : this.products[index].Sale);
      this.saleValue = ((old + add)-curPodSale);

      M.updateTextFields();

    },
    updateTotalBasic(cost) {

      this.componentKey += 1;
      var old;

      if (this.saleValue == null || this.saleValue == 0) {
        old = 0;
      } else {
        old = this.saleValue;
      }

      this.saleValue = (parseInt(old) + parseInt(cost));

      M.updateTextFields();

    },
    removeItem(tableName, index) {
      console.log(tableName);

      switch(tableName) {
            case "Product":
              this.products.splice(index, 1);           
              break;
            case "Campaign":
              this.campaigns.splice(index, 1);  
              break;
            default:
              M.toast({html: 'error', classes: 'rounded'});
          }
    },
    selectedSubItem(row) {
      
      var insName = row.Table.toLowerCase();
      M.Modal.getInstance(document.getElementById('modal-' + insName)).close();

      var newRow = {};

      switch(row.Table) {
            case "Product":
              newRow.Table = row.Table;
              newRow.UUID = row.UUID;
              newRow.DocID = row.DocID;
              newRow.Name = row.Name;
              newRow.Category = row.Category;
              newRow.Cost = row.Cost;
              newRow.Quantity = 1;
              newRow.Sale = row.Cost;
              this.products.push(newRow);
              this.updateTotalBasic(newRow.Cost);             
              break;
            case "Campaign":
              newRow.Table = row.Table;
              newRow.UUID = row.UUID;
              newRow.DocID = row.DocID;
              newRow.Name = row.Name;
              newRow.Successful = false;
              this.campaigns.push(newRow); 
              break;
            case "Account":
              if (this.accountSearch == "Account") {
                this.accountID = row.DocID;
              this.accountName = row.Name;
              } else {
                this.resellerID = row.DocID;
                this.resellerName = row.Name;
              }
              

              break;
            default:
              M.toast({html: 'error', classes: 'rounded'});
          }

          M.updateTextFields();
    },
    removeAccount() {
      if (this.accountSearch == "Account") {
                this.accountID = "";
              this.accountName = "";
              } else {
                this.resellerID = "";
                this.resellerName = "";
              }
      M.Modal.getInstance(document.getElementById('modal-account')).close();
    },
    getAccount(searchField) {
      var self = this;

      window.backend.readSubData('Account').then(result => { 
        
        var returned = JSON.parse(JSON.stringify(result));

        if (returned.Data === null) {
            self.titles = returned.Titles;
            self.accountData = [];
            M.toast({html: result.ReturnErr, classes: 'rounded'});

          } else {

            console.log('ere');
            console.log(searchField);

            self.titles = returned.Titles;
            self.accountData = returned.Data;
            self.accountSearch = searchField;

            var filt = [];
            var i;
            for (i = 0; i < returned.Titles.length; i++) {
              filt.push(returned.Titles[i].prop);
            }

            self.filters[0].prop = filt;
          
            M.Modal.getInstance(document.getElementById('modal-account')).open();
          
          }
      });
    },
    getProducts() {
      var self = this;

      window.backend.readSubData('Product').then(result => { 
        
        var returned = JSON.parse(JSON.stringify(result));

        if (returned.Data === null) {
            self.titles = returned.Titles;
            self.productData = [];
            M.toast({html: result.ReturnErr, classes: 'rounded'});

          } else {

            self.titles = returned.Titles;
            self.productData = returned.Data;

            var filt = [];
            var i;
            for (i = 0; i < returned.Titles.length; i++) {
              filt.push(returned.Titles[i].prop);
            }
            self.filters[0].prop = filt;
          
            M.Modal.getInstance(document.getElementById('modal-product')).open();
          
          }
      });
    },
    getCampaigns() {
      var self = this;

      window.backend.readSubData('Campaign').then(result => { 
        
        var returned = JSON.parse(JSON.stringify(result));

        if (returned.Data === null) {
            self.titles = returned.Titles;
            self.campaignData = [];
            M.toast({html: result.ReturnErr, classes: 'rounded'});

          } else {

            self.titles = returned.Titles;
            self.campaignData = returned.Data;

            var filt = [];
            var i;
            for (i = 0; i < returned.Titles.length; i++) {
              filt.push(returned.Titles[i].prop);
            }
            self.filters[0].prop = filt;
          
            M.Modal.getInstance(document.getElementById('modal-campaign')).open();
          
          }
      });
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
            this.description = result.Data.Description;
            this.sType = result.Data.SType;
            this.sSubType = result.Data.SSubType;
            this.accountID = result.Data.AccountID;
            this.accountName = result.Data.AccountName;
            this.status = result.Data.Status;
            this.stage = result.Data.Stage;
            this.leadsourceType = result.Data.LeadsourceType;
            this.leadsourceChannel = result.Data.LeadsourceChannel;
            this.saleValue = result.Data.SaleValue;
            this.estCloseDate = result.Data.EstCloseDate;
            this.actualCloseDate = result.Data.ActualCloseDate;
            this.closeReason = result.Data.CloseReason; 
            this.quoteID = result.Data.QuoteID; 
            this.invoiceID = result.Data.InvoiceID;
            this.docID = result.Data.DocID;
            this.createdDate = result.Data.CreatedDate;
            this.createdUsername = result.Data.CreatedUsername;
            this.activityCount = result.Data.ActivityCount;
            this.products = JSON.parse(result.Data.Products);
            this.campaigns = JSON.parse(result.Data.Campaigns);

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
      args.Description = self.description;
      args.SType = self.sType;
      args.SSubType = self.sSubType;
      args.AccountID = self.accountID;
      args.AccountName = self.accountName;
      // args.ResellerID = self.resellerID;
      // args.ResellerName = self.resellerName;
      args.Status = self.status;
      args.Stage = self.stage;
      args.LeadsourceType = self.leadsourceType;
      args.LeadsourceChannel = self.leadsourceChannel;
      args.SaleValue = self.saleValue;
      args.EstCloseDate = self.estCloseDate;
      args.ActualCloseDate = self.actualCloseDate;
      args.CloseReason = self.closeReason; 
      args.QuoteID = self.quoteID; 
      args.InvoiceID = self.invoiceID;
      args.Products = self.products;
      args.Campaigns = self.campaigns; 

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
      args.Description = self.description;
      args.SType = self.sType;
      args.SSubType = self.sSubType;
      args.AccountID = self.accountID;
      args.AccountName = self.accountName;
      // args.ResellerID = self.resellerID;
      // args.ResellerName = self.resellerName;
      args.Status = self.status;
      args.Stage = self.stage;
      args.LeadsourceType = self.leadsourceType;
      args.LeadsourceChannel = self.leadsourceChannel;
      args.SaleValue = self.saleValue;
      args.EstCloseDate = self.estCloseDate;
      args.ActualCloseDate = self.actualCloseDate;
      args.CloseReason = self.closeReason; 
      args.QuoteID = self.quoteID; 
      args.InvoiceID = self.invoiceID;
      args.Products = self.products;
      args.Campaigns = self.campaigns;
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
