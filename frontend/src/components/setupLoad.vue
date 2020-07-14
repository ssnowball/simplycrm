<template>
  <div class="container new-width">
    <h3 >Syncing up database</h3>
    <div class="preloader-wrapper big active">
    <div class="spinner-layer spinner-green-only">
      <div class="circle-clipper left">
        <div class="circle"></div>
      </div><div class="gap-patch">
        <div class="circle"></div>
      </div><div class="circle-clipper right">
        <div class="circle"></div>
      </div>
    </div>
  </div>
  <h5>It may take a few minutes</h5>
  </div>
</template>

<script>
import M from 'materialize-css'

export default {
  name: 'setupLoad',
  props: [],
  data() {
    return {
      // toLoad: "",
    };
  },
  mounted: function () {
    M.AutoInit();  
},
watch: {
  // status: function () {
  //     if (this.updateID) {
  //     var s = document.getElementById('fStatus');
  //     var v = this.status;
  //     for ( var i = 0; i < s.options.length; i++ ) {
  //       if ( s.options[i].text == v ) {
  //           s.options[i].selected = true;
  //           M.FormSelect.init(s);
  //           return;
  //       }
  //   }
  //     }
  //   },
},
  methods: {
    // deleteData: function() {
    //   document.getElementById("loader-show").classList.toggle('hide');

    //   var self = this;

    //   window.backend.deleteData(self.table, self.updateID).then(result => {

    //     if (result.Success === true) {
    //       M.toast({html: "Deleted Item", classes: 'rounded'});   
    //       this.changePage(); 
    //     } else {
    //       M.toast({html: result.ReturnErr, classes: 'rounded'});
    //     }

    //     document.getElementById("loader-show").classList.toggle('hide');
          
    //     });

    // },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
