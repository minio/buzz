$(document).ready(function () {

    // Issues Table
    var dataTableIssues = $('#table-issues').DataTable({
        "ajax": {
            "url": "/getIssues",
            "dataSrc": ""
        },
        lengthMenu: [[15, 30, 60, 100, -1], ['15 Rows', '30 Rows', '60 Rows', '100 Rows', 'All']],
        "sDom": '<"dataTables__top"lf><"dataTables__inner"rt><"dataTables__bottom"p><"clear">',
        "autoWidth": false,
        "columns": [
            { data : "number" },
            { data : "title" },
            {
                data : "Labels",
                "render": function(data) {
                    if(data !== null) {
                        var tag = '';
                        $.each(data, function(k, value) {
                            tag += "<div class='tableTables__tag' style='background-color:#"+value.color+"'>" + value.name + "</div>";
                        });
                        return tag;
                    }
                    else {
                        return "";
                    }
                }
            },
            { data : "login" },
            { data : "milestone"},
            { data : "state" },
            { data : "hours"},
            { data : "repository_url"}
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
            "url": "/getPRs",
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
            { data : "hours"},
            { data : "repo_name"}
        ]
    });


    // Table Column Resize
    $('.table').colResizable({
        liveDrag:true,
        resizeMode: 'flex'
    });


    // Issues table row link
    $('body').on('click', '#table-issues tbody tr', function () {
        var data = dataTableIssues.row( this ).data();
        window.open(data["html_url"], '_blank');
    });
});