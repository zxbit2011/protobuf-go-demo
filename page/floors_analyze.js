
/*
* 数据解析
* */

 function FloorsAnalyze(pbf, end) {

    return pbf.readFields((tag, el, pbf) => {
        if (tag === 1)
            el.floor.push(FeaturesAnalyze(pbf, pbf.readVarint() + pbf.pos));
        // else if (tag === 2)
        //     el.facility.push(FeaturesAnalyze(pbf, pbf.readVarint() + pbf.pos));
        else if (tag === 3)
            el.fill.push(FeaturesAnalyze(pbf, pbf.readVarint() + pbf.pos));
        else if (tag === 4)
            el.label.push(FeaturesAnalyze(pbf, pbf.readVarint() + pbf.pos));
    }, {
        floor: [],
        // facility: [],
        fill: [],
        label: []
    }, end);

}

// features
function FeaturesAnalyze(pbf, end) {
    return pbf.readFields((tag, el, pbf) => {
        if (tag === 1)
            el.geometry = GeometryAnalyze(pbf, pbf.readVarint() + pbf.pos);
        else if (tag === 2)
            el.properties = PropertiesAnalyze(pbf, pbf.readVarint() + pbf.pos);
    }, {
        geometry: null,
        properties: null
    }, end);
}

// geometry
function GeometryAnalyze(pbf, end) {
    return pbf.readFields((tag, el, pbf) => {
        if (tag === 1)
            el.type = pbf.readString();
        else if (tag === 2)
            el.coordinates = PointAnalyze(pbf, pbf.readVarint() + pbf.pos).point;
        else if (tag === 3)
            el.coordinates = CoordinatesAnalyze(pbf, pbf.readVarint() + pbf.pos).coordinates;
        else if (tag === 4)
            el.coordinates.push(CoordinatesAnalyze(pbf, pbf.readVarint() + pbf.pos).coordinates);
    }, {
        type: "",
        coordinates: []
    }, end);
}

// point
function PointAnalyze(pbf, end) {
    return pbf.readFields((tag, el, pbf) => {
        pbf.readPackedDouble(el.point);
    }, {
        point: []
    }, end);
}

// coordinates
function CoordinatesAnalyze(pbf, end) {
    return pbf.readFields((tag, el, pbf) => {
        el.coordinates.push(PointAnalyze(pbf, pbf.readVarint() + pbf.pos).point);
    }, {coordinates: []}, end);
}

// properties
function PropertiesAnalyze(pbf, end) {
    return pbf.readFields((tag, el, pbf) => {

        if (tag === 1) el["color"] = pbf.readString();
        else if (tag === 2) el["opacity"] = pbf.readFloat();
        else if (tag === 3) el["borderColor"] = pbf.readString();
        else if (tag === 6) el["id"] = pbf.readString();
        else if (tag === 7) el["name"] = pbf.readString();
        else if (tag === 9) el["x"] = pbf.readDouble();
        else if (tag === 10) el["y"] = pbf.readDouble();
        else if (tag === 11) el["floor"] = pbf.readVarint(true);
        else if (tag === 12) el["layer"] = pbf.readVarint();
        else if (tag === 15) el["height"] = pbf.readDouble();
        else if (tag === 16) el["base"] = pbf.readDouble();
        else if (tag === 18) el["icon"] = pbf.readString();

        // if (tag === 1) el["fill-color"] = pbf.readString();
        // else if (tag === 2) el["fill-opacity"] = pbf.readFloat();
        // else if (tag === 3) el["outline-color"] = pbf.readString();
        // else if (tag === 4) el["outline-width"] = pbf.readFloat();
        // else if (tag === 5) el["GEO_ID"] = pbf.readString();
        // else if (tag === 6) el["POI_ID"] = pbf.readString();
        // else if (tag === 7) el["NAME"] = pbf.readString();
        // else if (tag === 8) el["CATEGORY_ID"] = pbf.readString();
        // else if (tag === 9) el["LABEL_X"] = pbf.readDouble();
        // else if (tag === 10) el["LABEL_Y"] = pbf.readDouble();
        // else if (tag === 11) el["floor"] = pbf.readVarint(true);
        // else if (tag === 12) el["layer"] = pbf.readVarint();
        // else if (tag === 13) el["v"] = pbf.readVarint();
        // else if (tag === 14) el["extrusion"] = pbf.readBoolean();
        // else if (tag === 15) el["extrusion-height"] = pbf.readDouble();
        // else if (tag === 16) el["extrusion-base"] = pbf.readDouble();
        // else if (tag === 17) el["extrusion-opacity"] = pbf.readDouble();
        // else if (tag === 18){
        //     let image_name = pbf.readString();
        //     if(image_name){
        //         el["image-normal"] = image_name + "_normal";
        //     }
        // }
        // else if (tag === 19) el["ZONE_ID"] = pbf.readString();
        // else if (tag === 20) el["LEVEL_MIN"] = pbf.readVarint();
        // else if (tag === 21) el["LEVEL_MAX"] = pbf.readVarint();

    }, {
        "id": "",
        "name": "",
        "icon": "",
        "x": null,
        "y": null,
        "floor": null,
        "height": null,
        "base": null,
        "color": "",
        "opacity": ""

        // "fill-color": "",
        // "fill-opacity": '',
        // "outline-color": "",
        // "outline-width": '',
        // "GEO_ID": '',
        // "POI_ID": '',
        // "NAME": '',
        // "CATEGORY_ID": '',
        // "LABEL_X": null,
        // "LABEL_Y": null,
        // "floor": null,
        // "layer": null,
        // "v": '',
        // "extrusion": false,
        // "extrusion-height": null,
        // "extrusion-base": null,
        // "extrusion-opacity": null,
        // "image-normal": "",
        // "ZONE_ID": "",
        // "LEVEL_MIN": 0,
        // "LEVEL_MAX": 0
    }, end);
}
