"use strict";(self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[]).push([[195],{5882:function(t,i,e){e.d(i,{c:function(){return o}});const o=[{name:"复制链接",icon:"link"}]},6107:function(t,i,e){e.d(i,{Z:function(){return k}});var o=e(3396),n=e(9242),c=e(7139);const s=t=>((0,o.dD)("data-v-b546d348"),t=t(),(0,o.Cn)(),t),l={class:"topic-title"},a={class:"topic-intro van-multi-ellipsis--l2"},r=s((()=>(0,o._)("p",{class:"yes-ratio"},"Yes",-1))),u=s((()=>(0,o._)("p",{class:"no-ratio"},"No",-1)));function p(t,i,e,s,p,d){const m=(0,o.up)("van-loading"),f=(0,o.up)("van-image"),v=(0,o.up)("van-icon"),_=(0,o.up)("van-progress"),g=(0,o.up)("van-divider");return(0,o.wg)(),(0,o.iD)("div",{class:"single-topic",onClick:i[1]||(i[1]=i=>t.handleTopicClick(t.topic?.tid))},[(0,o.Wm)(f,{src:t.topic?.img_url,fit:"cover","lazy-load":"",alt:t.topic?.title,radius:"10px"},{loading:(0,o.w5)((()=>[(0,o.Wm)(m,{type:"spinner",size:"20"})])),error:(0,o.w5)((()=>[(0,o.Uk)("error")])),_:1},8,["src","alt"]),(0,o.Wm)(v,{name:"star",class:"icon-star",size:"2rem",color:!0===t.isCollect?"#fbff25":"#838383",onClick:(0,n.iM)(t.handleCollectClick,["stop"])},null,8,["color","onClick"]),(0,o._)("p",l,(0,c.zw)(t.topic?.title),1),(0,o._)("h4",a,(0,c.zw)(t.topic?.intro),1),r,u,(0,o.Wm)(_,{percentage:100*t.topic?.yes_ratio,"stroke-width":"16",color:"#61D089","pivot-color":t.topic?.yes_ratio<.5?"#f2a4a4":"#61D089"},null,8,["percentage","pivot-color"]),(0,o._)("div",null,[(0,o.Wm)(v,{name:"points",class:"icon-points",size:"1rem"},{default:(0,o.w5)((()=>[(0,o._)("span",null,(0,c.zw)(t.topic?.total_price),1)])),_:1}),(0,o.Wm)(v,{name:"guide-o",class:"icon-guide",size:"1rem",onClick:i[0]||(i[0]=(0,n.iM)((i=>t.handleShareClick(t.topic?.tid)),["stop"]))})]),(0,o.Wm)(g)])}e(7658);var d=e(7520),m=e(4870),f=e(2483),v=e(7330),_=e(9733),g=e(5882),w=(0,o.aZ)({props:{topic:{type:Object,required:!1}},emits:["showShareCard"],setup(t,{emit:i}){const e=(0,m.iH)(!1),o=(0,f.tv)(),n=(0,v.oR)(),c=(0,f.yj)(),s=n.state.login.userInfo,l=async()=>{if(n.state.login.userInfo.mixin_uuid){const i=await(0,d.oj)("/collect/check",{user_id:n.state.login.userInfo.mixin_uuid,tid:t.topic?.tid});null!==i.data&&(e.value=!0)}};l();const a=async()=>{if(s.mixin_uuid){if(!0===e.value){const i=await(0,d.oj)("/collect/delete",{user_id:n.state.login.userInfo.mixin_uuid,tid:t.topic?.tid});null!==i.data&&(e.value=!1),(0,_.F)("删除收藏!!!")}else{const i=await(0,d.oj)("/collect/add",{user_id:n.state.login.userInfo.mixin_uuid,tid:t.topic?.tid});null===i.data&&(e.value=!0),(0,_.F)("收藏成功!!!")}n.dispatch("main/handleCollecTopictAction")}else n.dispatch("login/accountLoginAction",c.path)},r=t=>{o.push(`/topic/${t}`)},u=t=>{console.log("handleShareClick"),i("showShareCard",t)};return{shareOptions:g.c,isCollect:e,handleCollectClick:a,handleTopicClick:r,handleShareClick:u}}}),h=e(89);const C=(0,h.Z)(w,[["render",p],["__scopeId","data-v-b546d348"]]);var k=C},4195:function(t,i,e){e.r(i),e.d(i,{default:function(){return T}});var o=e(3396),n=e(7139);const c={class:"user"},s={class:"userinfo-name"},l={class:"userinfo_identity"};function a(t,i,e,a,r,u){const p=(0,o.up)("van-image"),d=(0,o.up)("van-card"),m=(0,o.up)("user-buy"),f=(0,o.up)("van-tab"),v=(0,o.up)("user-collect"),_=(0,o.up)("van-tabs");return(0,o.wg)(),(0,o.iD)("div",c,[(0,o.Wm)(d,{centered:"","lazy-load":""},{thumb:(0,o.w5)((()=>[(0,o.Wm)(p,{class:"userinfo-avatar",src:t.userInfo.avatar_url,round:"",position:"left",width:"4rem",height:"4rem","lazy-load":""},null,8,["src"])])),title:(0,o.w5)((()=>[(0,o._)("p",s,(0,n.zw)(t.userInfo.full_name),1)])),desc:(0,o.w5)((()=>[(0,o._)("p",l,(0,n.zw)(t.userInfo.identity_number),1)])),_:1}),(0,o.Wm)(_,null,{default:(0,o.w5)((()=>[(0,o.Wm)(f,{title:t.$t("buy")},{default:(0,o.w5)((()=>[(0,o.Wm)(m)])),_:1},8,["title"]),(0,o.Wm)(f,{title:t.$t("collect")},{default:(0,o.w5)((()=>[(0,o.Wm)(v)])),_:1},8,["title"])])),_:1})])}var r=e(7330);const u={class:"user-buy"};function p(t,i,e,n,c,s){const l=(0,o.up)("single-topic");return(0,o.wg)(),(0,o.iD)("div",u,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(t.userToTopicList,(t=>((0,o.wg)(),(0,o.j4)(l,{key:t,topic:t.topic},null,8,["topic"])))),128))])}var d=e(6107),m=(0,o.aZ)({components:{SingleTopic:d.Z},setup(){const t=(0,r.oR)(),i=(0,o.Fl)((()=>t.state.main.userToTopicList));return{userToTopicList:i}}}),f=e(89);const v=(0,f.Z)(m,[["render",p],["__scopeId","data-v-9c394c5c"]]);var _=v;const g={class:"user-collect"};function w(t,i,e,n,c,s){const l=(0,o.up)("single-topic");return(0,o.wg)(),(0,o.iD)("div",g,[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(t.userCollectTopicList,(t=>((0,o.wg)(),(0,o.j4)(l,{key:t,topic:t.topic},null,8,["topic"])))),128))])}var h=(0,o.aZ)({components:{SingleTopic:d.Z},setup(){const t=(0,r.oR)(),i=(0,o.Fl)((()=>t.state.main.userCollectTopicList));return{userCollectTopicList:i}}});const C=(0,f.Z)(h,[["render",w],["__scopeId","data-v-217c7348"]]);var k=C,y=(0,o.aZ)({components:{UserBuy:_,UserCollect:k},setup(){const t=(0,r.oR)(),i=(0,o.Fl)((()=>t.state.login.userInfo));return{userInfo:i}}});const I=(0,f.Z)(y,[["render",a],["__scopeId","data-v-881640d8"]]);var T=I}}]);
//# sourceMappingURL=195.6fd0ea05.js.map