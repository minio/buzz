$(document).ready(function () {
    // Initiate data table
    $('.table--data').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getIssues",
            "dataSrc": ""
        },
        "sDom": '<"dataTables__top"lf>rt<"dataTables__bottom"p><"clear">',
        "columns": [
            { data : "number" },
            { data : "title" },
            { data : "name" },
            { data : "login" },
            { data : "milestone"},
            { data : "state" },
            { data : "repository_url"}
        ]
    });

    $('.table--data-1').DataTable({
        "ajax": {
            "url": "http://localhost:7000/getPRs",
            "dataSrc": ""
        },
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