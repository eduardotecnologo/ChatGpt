use std::os::raw::c_char;
use std::ffi::{CStr, CString};

#[no_mangle]
pub extern "C" fn hello_to_my_name(name: *const c_char) -> *mut c_char{
    let name = unsafe { CStr::from_ptr(name).to_str().unwrap() };
    let result = format!("Hello, {}!", name); 
    let result = CString::new(result).unwrap(); 
    result.into_raw()
}
