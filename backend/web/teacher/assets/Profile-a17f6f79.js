import{s as r,o,c as a,a as e,u as s,t as n,F as u,h as d,p,j as g}from"./teacher-6ecb1640.js";import{_ as h}from"./_plugin-vue_export-helper-c27b6911.js";import"./main-6cdd4766.js";const _=c=>(p("data-v-e06b1d5b"),c=c(),g(),c),m={class:"component"},f={class:"title"},b={class:"block-avatar"},v=["src"],I={class:"component-content"},k={class:"about-me"},y=_(()=>e("h3",null,"О себе",-1)),x={class:"language-items"},S=_(()=>e("h3",null,"Преподавание языков",-1)),B={key:0},F={key:1},L={__name:"Profile",setup(c){const t=r();return(P,A)=>(o(),a("div",m,[e("section",f,[e("figure",b,[e("img",{src:s(t).getAvatar,alt:"Аватар"},null,8,v),e("figcaption",null,[e("h1",null,n(s(t).getFullName),1),e("p",null,n(s(t).getTextIsTeacher),1),e("p",null,"Дата регистрации: "+n(s(t).getDate),1)])])]),e("section",I,[e("div",k,[y,e("pre",null,n(s(t).getAboutMe),1)]),e("div",x,[S,s(t).getLanguageItems.length!==0?(o(),a("ul",B,[(o(!0),a(u,null,d(s(t).getLanguageItems,(l,i)=>(o(),a("li",{key:`lang_item_${i}`},n(l),1))),128))])):(o(),a("p",F,"Вы не указали преподаваемые языки"))])])]))}},j=h(L,[["__scopeId","data-v-e06b1d5b"]]);export{j as default};
//# sourceMappingURL=Profile-a17f6f79.js.map
