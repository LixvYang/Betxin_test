"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[949],{8038:function(e,t,i){i.d(t,{Z:function(){return g}});var o=i(3396);const n={class:"page-topic"};function a(e,t,i,a,c,l){const s=(0,o.up)("single-topic"),r=(0,o.up)("van-list"),u=(0,o.up)("van-share-sheet");return(0,o.wg)(),(0,o.iD)("div",n,[(0,o.Wm)(r,{loading:e.loading,"onUpdate:loading":t[0]||(t[0]=t=>e.loading=t),finished:e.finished,"finished-text":"No More",onLoad:e.getTopicList},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.topicList,(t=>((0,o.wg)(),(0,o.j4)(s,{key:t,topic:t,onShowShareCard:e.showShareCard},null,8,["topic","onShowShareCard"])))),128))])),_:1},8,["loading","finished","onLoad"]),(0,o.Wm)(u,{show:e.showShare,"onUpdate:show":t[1]||(t[1]=t=>e.showShare=t),title:"Share",options:e.shareOptions,onSelect:e.onSelectShare},null,8,["show","options","onSelect"])])}var c=i(7330),l=i(4870),s=i(6610),r=i(7520),u=i(9733),p=i(5882),d=(0,o.aZ)({components:{SingleTopic:s.Z},props:{cid:{type:Number,required:!1},name:{type:String,required:!0}},setup(e){const t=(0,l.iH)(),i=(0,c.oR)(),n=(0,l.iH)(!1),a=(0,l.iH)(!1),s=(0,l.iH)(!1),d=(0,l.iH)(0),h=(0,o.Fl)((()=>i.getters["main/pageListData"](e.name))),m=(0,o.Fl)((()=>i.getters["main/pageListCount"](e.name))),v=()=>{let t={offset:h.value.length,limit:10};(0,r.bD)(`/topic/${e.cid}`,t).then((t=>{0===m.value||h.value.length<m.value?(i.commit(`main/changeTopic${e.name}Count`,t.data?.totalCount),i.commit(`main/pushTopic${e.name}List`,t.data?.list),a.value=!1,h.value?.length>=m.value&&(s.value=!0)):h.value.length===m.value&&(s.value=!0)}))},g=()=>{d.value=0,s.value=!1},f=e=>{(0,u.F)(e.name),n.value=!1;const i=document.createElement("input");document.body.appendChild(i),i.setAttribute("readonly","readonly"),i.setAttribute("value",`https://betxin.one/topic/${t.value}`),i.select(),document.execCommand("copy"),document.body.removeChild(i)},w=e=>{n.value=!0,t.value=e};return{showShare:n,topicList:h,finished:s,loading:a,getTopicList:v,changeFinished:g,onSelectShare:f,shareOptions:p.c,showShareCard:w}}}),h=i(89);const m=(0,h.Z)(d,[["render",a],["__scopeId","data-v-1e5a3e93"]]);var v=m,g=v},5882:function(e,t,i){i.d(t,{c:function(){return o}});const o=[{name:"复制链接",icon:"link"}]},5949:function(e,t,i){i.r(t),i.d(t,{default:function(){return p}});var o=i(3396);const n={class:"business"};function a(e,t,i,a,c,l){const s=(0,o.up)("page-topic");return(0,o.wg)(),(0,o.iD)("div",n,[(0,o.Wm)(s,{cid:3,name:"Sports"})])}var c=i(8038),l=(0,o.aZ)({components:{PageTopic:c.Z},setup(){return{}}}),s=i(89);const r=(0,s.Z)(l,[["render",a]]);var u=r,p=u},6610:function(e,t,i){i.d(t,{Z:function(){return y}});var o=i(3396),n=i(9242),a=i(7139);const c=e=>((0,o.dD)("data-v-7e33c823"),e=e(),(0,o.Cn)(),e),l={class:"topic-title"},s={class:"topic-intro van-multi-ellipsis--l2"},r=c((()=>(0,o._)("p",{class:"yes-ratio"},"Yes",-1))),u=c((()=>(0,o._)("p",{class:"no-ratio"},"No",-1))),p={class:"yes-ratio-price"},d={class:"no-ratio-price"};function h(e,t,i,c,h,m){const v=(0,o.up)("van-loading"),g=(0,o.up)("van-image"),f=(0,o.up)("van-icon"),w=(0,o.up)("van-progress"),C=(0,o.up)("van-divider");return(0,o.wg)(),(0,o.iD)("div",{class:"single-topic",onClick:t[1]||(t[1]=t=>e.handleTopicClick(e.topic?.tid))},[(0,o.Wm)(g,{src:e.topic?.img_url,fit:"cover","lazy-load":"",alt:e.topic?.title,radius:"10px"},{loading:(0,o.w5)((()=>[(0,o.Wm)(v,{type:"spinner",size:"20"})])),error:(0,o.w5)((()=>[(0,o.Uk)("error")])),_:1},8,["src","alt"]),(0,o.Wm)(f,{name:"star",class:"icon-star",size:"2rem",color:!0===e.isCollect?"#fbff25":"#838383",onClick:(0,n.iM)(e.handleCollectClick,["stop"])},null,8,["color","onClick"]),(0,o._)("p",l,(0,a.zw)(e.topic?.title),1),(0,o._)("h4",s,(0,a.zw)(e.topic?.intro),1),r,u,(0,o.Wm)(w,{percentage:100*e.topic?.yes_ratio,"stroke-width":"16",color:"#61D089","pivot-color":e.topic?.yes_ratio<.5?"#f2a4a4":"#61D089"},null,8,["percentage","pivot-color"]),(0,o._)("div",null,[(0,o.Wm)(f,{name:"points",class:"icon-points",size:"1rem"},{default:(0,o.w5)((()=>[(0,o._)("span",null,(0,a.zw)(e.topic?.total_price),1)])),_:1}),(0,o._)("span",p,"YES "+(0,a.zw)(e.topic?.yes_ratio_price),1),(0,o._)("span",d,"NO "+(0,a.zw)(e.topic?.no_ratio_price),1),(0,o.Wm)(f,{name:"guide-o",class:"icon-guide",size:"1rem",onClick:t[0]||(t[0]=(0,n.iM)((t=>e.handleShareClick(e.topic?.tid)),["stop"]))})]),(0,o.Wm)(C)])}i(7658);var m=i(7520),v=i(4870),g=i(2483),f=i(7330),w=i(9733),C=i(5882),_=(0,o.aZ)({props:{topic:{type:Object,required:!1}},emits:["showShareCard"],setup(e,{emit:t}){const i=(0,v.iH)(!1),o=(0,g.tv)(),n=(0,f.oR)(),a=(0,g.yj)(),c=n.state.login.userInfo,l=async()=>{if(n.state.login.userInfo.mixin_uuid){const t=await(0,m.oj)("/collect/check",{user_id:n.state.login.userInfo.mixin_uuid,tid:e.topic?.tid});null!==t.data&&(i.value=!0)}};l();const s=async()=>{if(c.mixin_uuid){if(!0===i.value){const t=await(0,m.oj)("/collect/delete",{user_id:n.state.login.userInfo.mixin_uuid,tid:e.topic?.tid});null!==t.data&&(i.value=!1),(0,w.F)("删除收藏!!!")}else{const t=await(0,m.oj)("/collect/add",{user_id:n.state.login.userInfo.mixin_uuid,tid:e.topic?.tid});null===t.data&&(i.value=!0),(0,w.F)("收藏成功!!!")}n.dispatch("main/handleCollecTopictAction")}else n.dispatch("login/accountLoginAction",a.path)},r=e=>{o.push(`/topic/${e}`)},u=e=>{console.log("handleShareClick"),t("showShareCard",e)};return{shareOptions:C.c,isCollect:i,handleCollectClick:s,handleTopicClick:r,handleShareClick:u}}}),S=i(89);const k=(0,S.Z)(_,[["render",h],["__scopeId","data-v-7e33c823"]]);var y=k}}]);
//# sourceMappingURL=949.94d44b30.js.map