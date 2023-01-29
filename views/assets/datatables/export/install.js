/* 
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */


apps.datatable_export = function (DataTable_Name, tag_location) {
    new $.fn.dataTable.Buttons(DataTable_Name, {
        buttons: [
            'copy', 'csv', 'excel', 'pdf', 'print'
        ]
    });
    $(tag_location).html(DataTable_Name.buttons().container());
};