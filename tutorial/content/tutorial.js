$(document).ready(init);

var max = 0;

function init() {
    var ie = ace.edit('iedit');
    ie.renderer.setShowGutter(false); 
    ie.setTheme('ace/theme/eclipse');
    ie.setHighlightActiveLine(false);
    ie.setShowPrintMargin(false);
    ie.getSession().setMode('ace/mode/unql');
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

    if ($('max').length > 0) max = $('#max').val();
    if (max < 1) max = guessmax();

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

    var slide = slideUrl(n);
	
    $.get(slide, function(data, status) {
        if (status != 'success') return;
        $('#content').html(data);

        var sample = $('#example').text();
        var ie = ace.edit('iedit');
        ie.setValue(sample);
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
    var query = 'q=' + encodeURIComponent(ie.getValue());
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
    var content = undefined;
    try
    {
        var json = $.parseJSON(data);
        for (var key in json) {
            if (json.hasOwnProperty(key) && key != 'resultset') {
                delete json[key];
            }
        }
        content = JSON.stringify(json, null, 2);
    }
    catch (e) {
        console.log(e);
        content = data;
    }

    var re = ace.edit('redit');
    re.setValue(content);
    re.navigateFileStart();
}

function slideUrl(n) {
    // note - .md is preprocessed into .html
    return 'slide-' + n + '.html';
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

function guessmax() {
    var http = new XMLHttpRequest();
    var n = 1;
    while (true) {
        try {
            http.open('HEAD', slideUrl(n), false);
            http.send();
            if (http.status >= 400) break;
            n = n + 1;
        }
        catch (e) {
            break;
        }
    }
    return n - 1;
}
