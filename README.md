Harsp
--

Harsp can quickly construct the Http Response structure, 
standardize the structure of the returned data, 
and currently supports five data formats: JSON, JSONP, XML, html, and text.

## 1. API Rsp
The returned data structure of the API is special. 
All API data is added with errCode and the corresponding message to indicate the result of this API request.

* JSON

    You can use this example when you want to return a successful JSON response data:
    
    ```
    harsp.JSON{Data:map[string]interface{}{"a":"b"}}.Success(w)
    ```
    This will return:
    
    > {"code":"0","data":{"a":"b"},"msg":"Success"}
    
    If you want to respond to a wrong data, then you can use this example:
    
    ```
    harsp.JSON{Rc:harsp.RetCode{Code:"10001", Msg:"err example"}}.Failed(w)
    ```
    This will return:
   
    > {"code":"10001","msg":"err example"}
    
    When returning an error response, you don't need Data to pass in. If you pass in Data, 
    Harsp will process it and return it together, but this is not the normal practice. 
    Since you think this is a bad request, why should you return data ?
    
* JSONP

    The data returned by JSONP is basically the same as the JSON data format, only the MIME type is modified, 
    and the function output of the callback is added.
    
    Success:
    
    ```
    harsp.JSONP{Data:map[string]interface{}{"a":"b"}, Callback:"Jsonp"}.Success(w)
    ```
    This will return:
    
    > Jsonp({"code":"0","data":{"a":"b"},"msg":"Success"})
    
    Failed:
    ```
    harsp.JSONP{Rc:harsp.RetCode{Code:"10002", Msg:"err example"}, Callback:"Jsonp"}.Failed(w)
    ```
    This will return:
    
    > Jsonp({"code":"10002","msg":"err example"})
    
* XML

    Harsp's API response data is defined as the map[string]interface{} structure type. 
    The official xml package does not support this type of xml conversion. 
    Therefore, Harsp applies the open source package of github.com/ryanwx/mxj to complete the API response. 
    Data map structure to xml conversion.

    Success:
    
    ```
    harsp.XML{Data:map[string]interface{}{"a":"b", "list":[]string{"aa", "bb"}}}.Success(w)
    ```
    This will return:
    
    ```
    <doc>
        <code>0</code>
        <data>
            <a>b</a>
            <list>aa</list>
            <list>bb</list>
        </data>
        <msg>Success</msg>
    </doc>
    ```
    
    Failed:
    ```
    harsp.XML{Rc:harsp.RetCode{Code:"10002", Msg:"err example"}}.Failed(w)
    ```
    This will return:
    
    ```
    <doc>
        <code>10002</code>
        <msg>err example</msg>
    </doc>
    ```
    
    In the xml conversion, the root tag is defined as doc. If it is a list structure, 
    the tag is defined as named tag, and the map is tagged by key.

* Return data default fill function

    ```
    DefaultPadding func(RetCode, interface{}) map[string]interface{} = defaultContentPadding
    
    func defaultContentPadding(rc RetCode, d interface{}) map[string]interface{} {
    	data := map[string]interface{}{
    		"code": rc.Code,
    		"msg":  rc.Msg,
    	}
    
    	if nil != d {
    		data["data"] = d
    	}
    
    	return data
    }
    ``` 
    Harsp defines a default return data fill function in the package. 
    If you want to customize the fill function, you can reset the DefaultPadding value when you start your application. 
    Once the settings are complete, Harsp returns the API class data. 
    The returned data will be populated with the func corresponding to this value
    
* Return error code structure

    ```
    type RetCode struct {
    	// return error code
    	// use this code to show this request result.
    	Code string
    	
    	// return error message
    	// this is request result message.
    	Msg  string
    }
    ```
    The return data structure of Harsp's API is contracted to return the code of the request structure 
    and the corresponding message unless Harsp returns a non-200 status code. Harsp specifies the structure of 
    the return data error code. You can customize the error code of their own application and pass in their 
    own defined error code when returning data.
    
* Default success response error code

    ```
    SuccessRet = RetCode{Code: "0", Msg: "Success"}
    ``` 
    When the server returns a response indicating success, no additional error code is required. 
    Harsp defines a default error code indicating success. Of course, you can reset the value.
    

## 2. Html OR text

The html OR Text data is returned without an error code, 
and Harsp will return the string completely without any padding.

* Html

    ```
    harsp.HTML{Data:"<h2>Hello World!</h2>"}.Send(w)
    ```    
    This will return:
        
    > ## Hello World!
    
* Text

    ```
    harsp.TEXT{Data:"<h2>Hello World!</h2>"}.Send(w)
    ```    
    This will return:
    
    ```
    <h2>Hello World!</h2>
    ``` 
    
<font color=red> All data return methods will return the http.ResponseWriter object of the current operation, 
so you can call the method to write data to the response body multiple times during the lifetime of your request, 
for html OR text data format The return is meaningful. If it is API data, 
Harsp recommends that all of it be written to the response data at once.</font>

Write response data multiple times:
--

Send:

```
harsp.HTML{Data:"<h2>Hello World!</h2>"}.Send(w)
harsp.TEXT{Data:"<h2>Hello World!</h2>"}.Send(w)
``` 

This will return:

> ## Hello World! <br/>Hello World!


License
--
Harsp is under the MIT license. See the LICENSE file for details.

Thinks
--
* xml transport open resource: [mxj](https://github.com/clbanning/mxj)