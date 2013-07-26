$(document).ready(init);

var max = undefined;

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

    max = $('#max').val();

    $('#run').click(run);
    $('#prev').click(prev).addClass('enabled');
    $('#next').click(next).addClass('enabled');
    if ('onhashchange' in window) $(window).bind('hashchange', change);

    load(getLocation());
}

function load(n) {
    if (n < 1 || n > max) return;

    setLocation(n);
    updateNav(n);

    var slide = slideUrl(n)
	var query = queryUrl(n)
	
    $.get(slide, function(data, status) {
        if (status != 'success') return;
        $('#content').html(data);
    });

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
    if (n == max-1) $('#next').removeClass('disabled').addClass('enabled');
    if (n == max) $('#next').removeClass('enabled').addClass('disabled');
}

function setLocation(n) {
    window.location.hash = '#' + n;
}

function getLocation(n) {
    var h = window.location.hash;
    if (!h || h.length < 2) return 1;
    var n = parseInt(h.substr(1));
    if (n >= 1 && n <= max) return n;
    return 1;
}

function next() {
    var n = getLocation();
    if (n < max) load(n + 1);
}

function prev() {
    var n = getLocation();
    if (n > 1) load(n - 1);
}

function change() {
    var n = getLocation();
    load (n);
}
