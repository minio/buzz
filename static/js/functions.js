$(document).ready(function () {

    // Data Table
    $('#table-issues').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getIssues",
            "dataSrc": ""
        },
        "initComplete": function(settings, json) {
            // ColResize
            $(this).colResizable({
                fixed:true,
                liveDrag:true,
                marginLeft: '5px'
            });
        },
        lengthMenu: [[15, 30, 60, 100, -1], ['15 Rows', '30 Rows', '60 Rows', '100 Rows', 'All']],
        "sDom": '<"dataTables__top"lf>rt<"dataTables__bottom"p><"clear">',
        "columns": [
            { data : "number" },
            { data : "title" },
            { data : "Labels[name]" },
            { data : "login" },
            { data : "milestone"},
            { data : "state" },
            { data : "repository_url"}
        ]
    });

    $('#table-prs').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getPRs",
            "dataSrc": ""
        },
        "initComplete": function(settings, json) {
            // ColResize
            $(this).colResizable({
                fixed:true,
                liveDrag:true,
                marginLeft: '5px'
            });
        },
        lengthMenu: [[15, 30, 60, 100, -1], ['15 Rows', '30 Rows', '60 Rows', '100 Rows', 'All']],
        "sDom": '<"dataTables__top"lf>rt<"dataTables__bottom"p><"clear">',
        "columns": [
            { data : "number" },
            { data : "title" },
            { data : "created_at" },
            { data : "updated_at"},
            { data : "repo_name"}
        ]
    });
});