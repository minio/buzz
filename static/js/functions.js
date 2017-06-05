$(document).ready(function () {
    var table;
    var tableData;
    var tableDataID;
    var tableDataRepo;

    // Notification
    function notify (title, text, type) {
        new PNotify({
            title: title,
            text: text,
            type: type,
            styling: 'bootstrap3',
            buttons: {
                closer: false,
                sticker: false
            },
            delay: 2500,
            animate: {
                animate: true,
                in_class: 'fadeInDown',
                out_class: 'fadeOutUp'
            }
        });
    }

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
                "targets": [ 8 ],
                "data": null,
                "defaultContent": "<div class='set-eta'><input class='set-eta__input' type='text' placeholder='Set ETA' /></div>"
            }
        ],
        "initComplete": function (settings, json) {

        },
        "rowCallback": function( row, data, index ) {
            if ( data[4] == "A" ) {
                $('td:eq(4)', row).html( '<b>A</b>' );
            }
        }
    });

    // PR Table
    var dataTablePR = $('#table-prs').DataTable({
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
            { data : "sender"},
            { data : "updated_at"},
            { data : "hours"},
            { data : "repo_name"},
            {
                data : "Reviewers",
                "render": function(data) {
                    if(data !== null) {
                        var reviewer = '';
                        $.each(data, function(k, value) {
                            reviewer += "<div class='tableTables__tag tableTables__tag--default'>" + value.user.login + "</div>";
                        });
                        return reviewer;
                    }
                    else {
                        return "";
                    }
                }
            },
            {
                data : "ReviewState",
                "render": function(data) {
                    if(data !== null) {
                        var reviewActivity = '';
                        $.each(data, function(k, value) {
                            reviewActivity += "<div class='tableTables__tag tableTables__tag--default'>" + value.user.login + " : "+ value.state + "</div>";
                        });
                        return reviewActivity;
                    }
                    else {
                        return "";
                    }
                }
            }
        ]
    });


    // Refresh data tables
    setInterval(function () {
        if($('#table-issues')[0]) {
            dataTableIssues.ajax.reload();
        }

        if($('#table-prs')[0]) {
            dataTablePR.ajax.reload();
        }
    }, 300000);


    // Table Column Resize
    $('.table').colResizable({
        liveDrag:true,
        resizeMode: 'flex'
    });

    // Table row link and data assign
    $('body').on('click', 'table tbody tr', function (e) {

        if($(this).closest('table').is('#table-prs')) {
            table = dataTablePR
        }

        if($(this).closest('table').is('#table-issues')) {
            table = dataTableIssues
        }

        // Get table data and assign to variables
        tableData = table.row(this).data();
        tableDataID = tableData["number"];
        var tableDataURL = tableData["repository_url"];
        tableDataRepo = tableDataURL.substr(tableData["repository_url"].lastIndexOf('/') + 1);

        // Open Issue link in new tab
        if(!$(e.target).is('.set-eta__input')) {
            window.open(tableData["html_url"], '_blank');
        }
    });

    $('body').on('click', '.set-eta__input', function () {
        var inputETA = $(this);

        // Add/remove active class to identify active instance
        //$(this).addClass('set-eta__input--active');

        // Remove any previous instances
        $('.flatpickr-calendar').remove();

        // Get current date
        var defaultDate = $(this).val() || '';

        // Initiate the picker
        $(this).flatpickr({
            defaultDate: defaultDate,
            enableTime: true,
            nextArrow: '<i class="zmdi zmdi-long-arrow-right" />',
            prevArrow: '<i class="zmdi zmdi-long-arrow-left" />',
            onClose: function(dateObj, dateStr) {
                if(dateStr != '') {
                    if(!inputETA.hasClass('set-eta__input--active')) {
                       $.post( '/setETA?number='+tableDataID+'&repo='+tableDataRepo+'&org=minio&comment=ETA: '+dateStr, function() {
                            notify('', 'ETA successfully set for issue no. '+ tableDataID + ' of ' + tableDataRepo + ' repository!', 'success');
                            inputETA.addClass('set-eta__input--active');
                       })
                    }
                }
            }
        }).open();

    });
});