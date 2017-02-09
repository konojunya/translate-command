#!/usr/bin/env node
'use strict';

const http = require('http');
const https = require('https');
const qs = require('querystring');

function getAccessToken(callback) {
	let body = '';
	let data = {
		'client_id': process.env.MS_TRANSLATE_ID,
		'client_secret': process.env.MS_TRANSLATE_SECRET,
		'scope': 'http://api.microsofttranslator.com',
		'grant_type': 'client_credentials'
	};

	let req = https.request({
		host: 'datamarket.accesscontrol.windows.net',
		path: '/v2/OAuth2-13',
		method: 'POST'
	}, (res) => {
		res.setEncoding('utf8');
		res.on('data', (chunk) => {
			body += chunk;
		}).on('end', () => {
			let resData = JSON.parse(body);
			callback(resData.access_token);
		});
	}).on('error', (err) => {
		console.log(err);
	});
	req.write(qs.stringify(data));
	req.end();
}

function translate(token, text,from,to, callback) {
	let options = 'from='+from+'&to='+to+'&text=' + qs.escape(text) +'&oncomplete=translated';
	let body = '';
	let req = http.request({
		host: 'api.microsofttranslator.com',
		path: '/V2/Ajax.svc/Translate?' + options,
		method: 'GET',
		headers: {
			"Authorization": 'Bearer ' + token
		}
	}, (res) => {
		res.setEncoding('utf8');
		res.on('data', (chunk) => {
			body += chunk;
		}).on('end', () => {
			eval(body);
		});
	}).on('error', (err) => {
		console.log(err);
	});
	req.end();

	function translated(text) {
		callback(text);
	}
}

var text = process.argv[2]
var translate_type = process.argv[3] || "ja/en"

translate_type = translate_type.split("/");

getAccessToken((token) => {
    translate(token, text,translate_type[0],translate_type[1], (translated) => {
        console.log(translated);
    });
});
