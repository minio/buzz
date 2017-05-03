$(document).ready(function () {

    // Issues
    var dataTableIssues = $('#table-issues').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getIssues",
            "dataSrc": ""
        },
        lengthMenu: [[15, 30, 60, 100, -1], ['15 Rows', '30 Rows', '60 Rows', '100 Rows', 'All']],
        "sDom": '<"dataTables__top"lf><"dataTables__inner"rt><"dataTables__bottom"p><"clear">',
        "autoWidth": false,
        "columns": [
            { data : "number" },
            { data : "title" },
            { data : "Labels[name]" },
            { data : "login" },
            { data : "milestone"},
            { data : "state" },
            { data : "repository_url"},
            { data : "html_url"}
        ],
        "columnDefs": [
            {
                "targets": [ -1 ],
                "visible": false,
                "searchable": false
            }
        ]
    });

    // PR Table
    $('#table-prs').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getPRs",
            "dataSrc": ""
        },
        "autoWidth": false,
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


    // Table Column Resize
    $('.table').colResizable({
        liveDrag:true,
        resizeMode: 'flex'
    });
});