let json;   //定义了一个全局变量，用来存储json文件,后面在传值的过程中怎么也得不到值，只好定义一个全局变量
$(()=>{

    $.ajax({
        url:"/page_soft/ListDevTools",
        type:"get",
        success: response=> {
            let data = JSON.parse(response)
            json = data
            for (i=0;i<data.length;i++) {
                let title = data[i]["SoftWareTitle"]
                let img = data[i]["SoftWareImg"]
                $("#box").prepend( ()=> {
                    return '<div class="layui-col-md2 click-block">' +
                                '<div class="layui-card" onclick="popup(&quot;'+i+'&quot;)" >'  +   //这里传了一个索引，是因为我尝试传送值失败
                                    '<div class="layui-card-body" style="height: 250px;">' +
                                        '<ul>' +
                                            '<li><img height="180" src="'+img+'"/></li>'+
                                            '<li>'+title+'</li>' +
                                        '</ul>'+
                                    '</div>'+
                                '</div>'+
                            '</div>'
                })

            }

        }
    })
})
function popup(index) {
    let title = json[index]["SoftWareTitle"]
    let desc = json[index]["SoftWareDescription"]
    let url = json[index]["SoftWareUrl"]
    layer.open({
        title: title
        ,content: desc
        ,btn:'立即下载'
        ,yes: function(){
            window.open(url)
        }
    })
}
