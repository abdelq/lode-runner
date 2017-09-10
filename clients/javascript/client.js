const { createConnection } = require('net');
const { Transform } = require('stream');

class LinerTransform extends Transform {
    _transform(chunk, encoding, callback) {
        chunk = chunk.toString()
        if (this._lastLineData) chunk = this._lastLineData + chunk

        var lines = chunk.split('\n')
        this._lastLineData = lines.splice(lines.length-1,1)[0]

        lines.forEach(this.push.bind(this))
        callback()
    }

    _flush(callback) {
        if (this._lastLineData) this.push(this._lastLineData)
        this._lastLineData = null
        callback()
    }
}

const client = createConnection(1337, () => {
});

client.pipe(new LinerTransform()).on('data', (data) => {
    console.log(data);
});
