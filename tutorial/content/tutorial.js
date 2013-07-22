$(document).ready(init);

var maxpage = undefined;

function init() {
    var ie = ace.edit('iedit');
    ie.renderer.setShowGutter(false); 
    ie.setTheme('ace/theme/eclipse');
    ie.setHighlightActiveLine(false);
    ie.setShowPrintMargin(false);
    ie.getSession().setMode('ace/mode/sql');
    ie.getSession().setUseWrapMode(true);
    ie.setDisplayIndentGuides(false);

    var re = ace.edit('redit');
    re.renderer.setShowGutter(false); 
    re.setTheme('ace/theme/eclipse');
    re.setHighlightActiveLine(false);
    re.setShowPrintMargin(false);
    re.getSession().setMode('ace/mode/json');
    re.getSession().setUseWrapMode(false);
    re.setReadOnly(false);
    re.setDisplayIndentGuides(true);
    re.setShowFoldWidgets(true);

    maxpage = $('#max').val();

    $('#run').click(run);
    $('#prev').click(prev).addClass('enabled');
    $('#next').click(next).addClass('enabled');

    load(1);
}

function load(n) {
    if (n < 1 || n > maxpage) return;

    page = n;
    updateNav(n);

    var slide = slideUrl(n)
	var query = queryUrl(n)
	
    $('#content').attr('src', slide);

	$.get(query, function(data, status) {
		if (status != 'success') return;
	    var ie = ace.edit('iedit');
	    ie.setValue(data);
	    ie.navigateFileEnd();
	
	    var re = ace.edit('redit');
	    re.setValue("  ");
	    re.navigateFileEnd();

	    ie.focus();
    });
}

function run() {
	var url = '/query';
    var ie = ace.edit('iedit');
    var query = 'q=' + ie.getValue();
    $.post(url, query, ran).fail(failed);
}

function failed(data) {
    var msg = data.status + ': ' + data.statusText + '\n\n';
    msg += data.responseText + '\n';
    var re = ace.edit('redit');
    re.setValue(msg);
    re.navigateFileStart();
}

function ran(data) {
    console.log("Ran");
    console.log(data);
    var re = ace.edit('redit');
    re.setValue(data);
    re.navigateFileStart();
}

function slideUrl(n) {
    // note - .md is preprocessed into .html
    return 'slide-' + n + '.html';
}

function queryUrl(n) {
    return 'query-' + n + '.unql';
}

function updateNav(n) {
    if (n == 1) $('#prev').removeClass('enabled').addClass('disabled');
    if (n == 2) $('#prev').removeClass('disabled').addClass('enabled');
    if (n == maxpage-1) $('#next').removeClass('disabled').addClass('enabled');
    if (n == maxpage) $('#next').removeClass('enabled').addClass('disabled');
}

var page = 0;

function next() {
    if (page == maxpage) return;
    load(page + 1);
}

function prev() {
    if (page == 1) return;
    load(page - 1);
}
