(this.webpackJsonpfrontend=this.webpackJsonpfrontend||[]).push([[0],{60:function(e,t,a){e.exports=a(91)},65:function(e,t,a){},66:function(e,t,a){},91:function(e,t,a){"use strict";a.r(t);var n=a(0),l=a.n(n),r=a(25),o=a.n(r),c=(a(65),a(58)),s=a(5),i=a(29),d=a(38),u=a(8),m=a(33),h=a(34),p=(a(66),a(19)),b=a(13),E=a(18),f=a(22),v=a(21),g=a(56),w=a(31),y=a.n(w),S=function(e){Object(f.a)(a,e);var t=Object(v.a)(a);function a(e){var n;return Object(p.a)(this,a),(n=t.call(this,e)).readData(),n.state={products:[]},n.readData=n.readData.bind(Object(b.a)(n)),n}return Object(E.a)(a,[{key:"readData",value:function(){var e=this;y.a.get(window.global.api_location+"/products").then((function(t){console.log(t.data),e.setState({products:t.data})})).catch((function(e){console.log(e)}))}},{key:"getProducts",value:function(){for(var e=[],t=0;t<this.state.products.length;t++)e.push(l.a.createElement("tr",{key:t},l.a.createElement("td",null,this.state.products[t].name),l.a.createElement("td",null,this.state.products[t].price),l.a.createElement("td",null,this.state.products[t].sku)));return e}}]),Object(E.a)(a,[{key:"render",value:function(){return l.a.createElement("div",null,l.a.createElement("h1",{style:{marginBottom:"40px"}},"Menu"),l.a.createElement(g.a,null,l.a.createElement("thead",null,l.a.createElement("tr",null,l.a.createElement("th",null,"Name"),l.a.createElement("th",null,"Price"),l.a.createElement("th",null,"SKU"))),l.a.createElement("tbody",null,this.getProducts())))}}]),a}(l.a.Component),j=a(28),O=a(26),k=a(43),x=a(57),D=a(39),C=function(e){Object(f.a)(a,e);var t=Object(v.a)(a);function a(e){var n;return Object(p.a)(this,a),(n=t.call(this,e)).state={show:!1},n.hide=n.hide.bind(Object(b.a)(n)),n}return Object(E.a)(a,[{key:"hide",value:function(){this.setState({show:!1})}},{key:"componentDidUpdate",value:function(e){this.props!==e&&this.setState({show:this.props.show})}},{key:"render",value:function(){return l.a.createElement(D.a,{onClose:this.hide,show:this.state.show,delay:3e3,autohide:!0},l.a.createElement(D.a.Header,null,l.a.createElement("strong",{className:"mr-auto"},"File Upload")),l.a.createElement(D.a.Body,null,this.props.message))}}]),a}(l.a.Component),T=function(e){Object(f.a)(a,e);var t=Object(v.a)(a);function a(e){var n;return Object(p.a)(this,a),(n=t.call(this,e)).state={validated:!1,id:"",buttonDisabled:!1,toastShow:!1,toastText:"asd"},n.validated=n.validated.bind(Object(b.a)(n)),n.changeHandler=n.changeHandler.bind(Object(b.a)(n)),n.handleSubmit=n.handleSubmit.bind(Object(b.a)(n)),n}return Object(E.a)(a,[{key:"validated",value:function(){return console.log("validated",this.state.validated),this.state.validated}},{key:"handleSubmit",value:function(e){var t=this;if(e.preventDefault(),!1!==e.currentTarget.checkValidity()){this.setState({buttonDisabled:!0,toastShow:!1});var a=new FormData;a.append("file",this.state.file),a.append("id",this.state.id),y.a.post(window.global.files_location,a,{"content-type":"multipart/form-data; boundary=".concat(a._boundary)}).then((function(e){console.log(e);var a="";a=200===e.status?"Uploaded file":"Unable to upload file. Error:"+e.statusText,t.setState({buttonDisabled:!1,toastShow:!0,toastText:a})})).catch((function(e){console.log("Err"+e),t.setState({buttonDisabled:!1,toastShow:!0,toastText:"Unable to upload file. "+e})}))}else e.stopPropagation()}},{key:"changeHandler",value:function(e){var t,a;"file"!==e.target.name?this.setState((t={},Object(j.a)(t,e.target.name,e.target.value),Object(j.a)(t,"toastShow",!1),t)):this.setState((a={},Object(j.a)(a,e.target.name,e.target.files[0]),Object(j.a)(a,"toastShow",!1),a))}},{key:"render",value:function(){return l.a.createElement("div",null,l.a.createElement("h1",{style:{marginBottom:"40px"}},"Admin"),l.a.createElement(x.a,{className:"text-left"},l.a.createElement(u.a,{noValidate:!0,validated:this.validated,onSubmit:this.handleSubmit},l.a.createElement(u.a.Group,{as:k.a,controlId:"productID"},l.a.createElement(u.a.Label,{column:!0,sm:"2"},"Product ID:"),l.a.createElement(O.a,{sm:"6"},l.a.createElement(u.a.Control,{type:"text",name:"id",placeholder:"",required:!0,style:{width:"80px"},value:this.state.id,onChange:this.changeHandler}),l.a.createElement(u.a.Text,{className:"text-muted"},"Enter the product id to upload an image for"),l.a.createElement(u.a.Control.Feedback,{type:"invalid"},"Please provide a product ID.")),l.a.createElement(O.a,{sm:"4"},l.a.createElement(C,{show:this.state.toastShow,message:this.state.toastText}))),l.a.createElement(u.a.Group,{as:k.a},l.a.createElement(u.a.Label,{column:!0,sm:"2"},"File:"),l.a.createElement(O.a,{sm:"10"},l.a.createElement(u.a.Control,{type:"file",name:"file",placeholder:"",required:!0,onChange:this.changeHandler}),l.a.createElement(u.a.Text,{className:"text-muted"},"Image to associate with the product"),l.a.createElement(u.a.Control.Feedback,{type:"invalid"},"Please select a file to upload."))),l.a.createElement(h.a,{type:"submit",disabled:this.state.buttonDisabled},"Submit form"))))}}]),a}(l.a.Component);a(87);var N=function(){return l.a.createElement(c.a,null,l.a.createElement("div",{className:"App"},l.a.createElement(i.a,{bg:"light",expand:"lg"},l.a.createElement(i.a.Brand,{href:"#home"},"Coffee Shop"),l.a.createElement(i.a.Toggle,{"aria-controls":"basic-navbar-nav"}),l.a.createElement(i.a.Collapse,{id:"basic-navbar-nav"},l.a.createElement(d.a,{className:"mr-auto"},l.a.createElement(d.a.Link,{href:"/"},"Home"),l.a.createElement(d.a.Link,{href:"/admin"},"Admin")),l.a.createElement(u.a,{inline:!0},l.a.createElement(m.a,{type:"text",placeholder:"Search",className:"mr-sm-2"}),l.a.createElement(h.a,{variant:"outline-success"},"Search")))),l.a.createElement(s.c,null,l.a.createElement(s.a,{path:"/admin"},l.a.createElement(T,null)),l.a.createElement(s.a,{path:"/"},l.a.createElement(S,null)),l.a.createElement(s.a,null,l.a.createElement(S,null)))))};Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));o.a.render(l.a.createElement(N,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))}},[[60,1,2]]]);
//# sourceMappingURL=main.af67dbe0.chunk.js.map