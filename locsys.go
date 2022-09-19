package autocpp

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"runtime"
	"strings"
)

type LocalSystem struct {
	commonIncludes           []string
	systemIncludeDirectories []string // can be searched
	localIncludeDirectories  []string // should not be searched exhaustively, because this slice includes ".."
	includeFiles             []string
	verbose                  bool
}

func (locsys *LocalSystem) SystemIncludeDirectories() []string {
	xs := make([]string, 0)
	switch runtime.GOOS {
	case "linux":
		xs = append(xs, "/usr/include")
	case "windows":
		// break
	default:
		xs = append(xs, "/usr/include")
	}
	return xs
}

func (locsys *LocalSystem) CommonIncludes() []string {
	// Skip C99, C++, C++20 and deprecated C++ headers + more
	xs := []string{"algorithm", "any", "array", "assert.h", "atomic", "barrier", "bit", "bitset", "cassert", "cctype", "cerrno", "cfenv", "cfloat", "charconv", "chrono", "cinttypes", "climits", "clocale", "cmath", "codecvt", "compare", "complex", "complex.h", "concepts", "condition_variable", "coroutine", "csetjmp", "csignal", "cstdarg", "cstddef", "cstdint", "cstdio", "cstdlib", "cstring", "ctime", "ctype.h", "deque", "errno.h", "exception", "execution", "fenv.h", "filesystem", "float.h", "format", "forward_list", "fstream", "functional", "future", "glibc", "initializer_list", "inttypes.h", "iomanip", "ios", "iosfwd", "iostream", "iso646.h", "istream", "iterator", "latch", "limits", "limits.h", "list", "locale", "locale.h", "map", "math.h", "memory", "memory_resource", "mutex", "new", "numbers", "numeric", "optional", "ostream", "queue", "random", "ranges", "ratio", "regex", "scoped_allocator", "semaphore", "set", "setjmp.h", "shared_mutex", "signal.h", "source_location", "span", "sstream", "stack", "stacktrace", "stdalign.h", "stdarg.h", "stdatomic.h", "stdbool.h", "stddef.h", "stdexcept", "stdint.h", "stdio.h", "stdlib.h", "stdnoreturn.h", "stop_token", "streambuf", "string", "string.h", "strstream", "syncstream", "system_error", "tgmath.h", "thread", "threads.h", "time.h", "tuple", "type_traits", "typeindex", "typeinfo", "uchar.h", "unordered_map", "unordered_set", "utility", "valarray", "variant", "vector", "version", "wchar.h", "wctype.h", "windows.h"}
	switch runtime.GOOS {
	case "windows":
		xs = append(xs, []string{"GL/gl.h", "GL/glaux.h", "GL/glcorearb.h", "GL/glext.h", "GL/glu.h", "GL/glxext.h", "GL/wglext.h", "accctrl.h", "aclapi.h", "aclui.h", "activation.h", "activaut.h", "activdbg.h", "activdbg100.h", "activecf.h", "activeds.h", "activprof.h", "activscp.h", "adc.h", "adhoc.h", "admex.h", "adoctint.h", "adodef.h", "adogpool.h", "adogpool_backcompat.h", "adoguids.h", "adoid.h", "adoint.h", "adoint_backcompat.h", "adojet.h", "adomd.h", "adptif.h", "adsdb.h", "adserr.h", "adshlp.h", "adsiid.h", "adsnms.h", "adsprop.h", "adssts.h", "adtgen.h", "advpub.h", "afxres.h", "af_irda.h", "agtctl.h", "agtctl_i.c", "agterr.h", "agtsvr.h", "agtsvr_i.c", "alg.h", "alink.h", "amaudio.h", "amstream.h", "amstream.idl", "amvideo.h", "amvideo.idl", "apdevpkey.h", "apiset.h", "apisetcconv.h", "appmgmt.h", "aqadmtyp.h", "asptlb.h", "assert.h", "atacct.h", "atalkwsh.h", "atsmedia.h", "audevcod.h", "audioapotypes.h", "audioclient.h", "audioendpoints.h", "audioengineendpoint.h", "audiopolicy.h", "audiosessiontypes.h", "austream.h", "austream.idl", "authif.h", "authz.h", "aux_ulib.h", "avifmt.h", "aviriff.h", "avrfsdk.h", "avrt.h", "axextendenums.h", "azroles.h", "basetsd.h", "basetyps.h", "batclass.h", "bcrypt.h", "bdaiface.h", "bdaiface_enums.h", "bdamedia.h", "bdatypes.h", "bemapiset.h", "bh.h", "bidispl.h", "bits.h", "bits1_5.h", "bits2_0.h", "bitscfg.h", "bitsmsg.h", "blberr.h", "bluetoothapis.h", "bthdef.h", "bthsdpdef.h", "bugcodes.h", "callobj.h", "cardmod.h", "casetup.h", "cchannel.h", "cderr.h", "cdoex.h", "cdoexerr.h", "cdoexm.h", "cdoexm_i.c", "cdoexstr.h", "cdoex_i.c", "cdonts.h", "cdosys.h", "cdosyserr.h", "cdosysstr.h", "cdosys_i.c", "celib.h", "certadm.h", "certbase.h", "certbcli.h", "certcli.h", "certenc.h", "certenroll.h", "certexit.h", "certif.h", "certmod.h", "certpol.h", "certreqd.h", "certsrv.h", "certview.h", "cfg.h", "cfgmgr32.h", "cguid.h", "chanmgr.h", "cierror.h", "clfs.h", "clfsmgmt.h", "clfsmgmtw32.h", "clfsw32.h", "cluadmex.h", "clusapi.h", "cluscfgguids.h", "cluscfgserver.h", "cluscfgwizard.h", "cmdtree.h", "cmnquery.h", "codecapi.h", "color.dlg", "colordlg.h", "comadmin.h", "combaseapi.h", "comcat.h", "comdef.h", "comdefsp.h", "comip.h", "comlite.h", "commapi.h", "commctrl.h", "commctrl.rh", "commdlg.h", "common.ver", "commoncontrols.h", "complex.h", "compobj.h", "compressapi.h", "compstui.h", "comsvcs.h", "comutil.h", "confpriv.h", "conio.h", "control.h", "cor.h", "corerror.h", "corhdr.h", "correg.h", "cpl.h", "cplext.h", "credssp.h", "crtdbg.h", "crtdefs.h", "cryptuiapi.h", "cryptxml.h", "cscapi.h", "cscobj.h", "ctfutb.h", "ctxtcall.h", "ctype.h", "custcntl.h", "d2d1.h", "d2d1effectauthor.h", "d2d1effecthelpers.h", "d2d1effects.h", "d2d1helper.h", "d2d1_1.h", "d2d1_1helper.h", "d2dbasetypes.h", "d2derr.h", "d3d.h", "d3d8.h", "d3d8caps.h", "d3d8types.h", "d3d9.h", "d3d9caps.h", "d3d9types.h", "d3d10.h", "d3d10.idl", "d3d10effect.h", "d3d10misc.h", "d3d10shader.h", "d3d10_1.h", "d3d10_1.idl", "d3d10_1shader.h", "d3d11.h", "d3d11.idl", "d3d11sdklayers.h", "d3d11sdklayers.idl", "d3d11shader.h", "d3d11_1.h", "d3d11_1.idl", "d3dcaps.h", "d3dcommon.h", "d3dcommon.idl", "d3dcompiler.h", "d3dhal.h", "d3drm.h", "d3drmdef.h", "d3drmobj.h", "d3dtypes.h", "d3dvec.inl", "d3dx9.h", "d3dx9anim.h", "d3dx9core.h", "d3dx9effect.h", "d3dx9math.h", "d3dx9math.inl", "d3dx9mesh.h", "d3dx9shader.h", "d3dx9shape.h", "d3dx9tex.h", "d3dx9xof.h", "daogetrw.h", "datapath.h", "datetimeapi.h", "davclnt.h", "dbdaoerr.h", "dbdaoid.h", "dbdaoint.h", "dbgautoattach.h", "dbgeng.h", "dbghelp.h", "dbgprop.h", "dbt.h", "dciddi.h", "dciman.h", "dcommon.h", "dcomp.h", "dcompanimation.h", "dcomptypes.h", "dde.h", "dde.rh", "ddeml.h", "ddk/acpiioct.h", "ddk/afilter.h", "ddk/amtvuids.h", "ddk/atm.h", "ddk/bdasup.h", "ddk/classpnp.h", "ddk/csq.h", "ddk/d3dhal.h", "ddk/d3dhalex.h", "ddk/d4drvif.h", "ddk/d4iface.h", "ddk/dderror.h", "ddk/dmusicks.h", "ddk/drivinit.h", "ddk/drmk.h", "ddk/dxapi.h", "ddk/fltsafe.h", "ddk/hidclass.h", "ddk/hubbusif.h", "ddk/ide.h", "ddk/ioaccess.h", "ddk/kbdmou.h", "ddk/mcd.h", "ddk/mce.h", "ddk/miniport.h", "ddk/minitape.h", "ddk/mountdev.h", "ddk/mountmgr.h", "ddk/msports.h", "ddk/ndis.h", "ddk/ndisguid.h", "ddk/ndistapi.h", "ddk/ndiswan.h", "ddk/netpnp.h", "ddk/ntagp.h", "ddk/ntddk.h", "ddk/ntddpcm.h", "ddk/ntddsnd.h", "ddk/ntifs.h", "ddk/ntimage.h", "ddk/ntnls.h", "ddk/ntpoapi.h", "ddk/ntstrsafe.h", "ddk/oprghdlr.h", "ddk/parallel.h", "ddk/pfhook.h", "ddk/poclass.h", "ddk/portcls.h", "ddk/punknown.h", "ddk/scsi.h", "ddk/scsiscan.h", "ddk/scsiwmi.h", "ddk/smbus.h", "ddk/srb.h", "ddk/stdunk.h", "ddk/storport.h", "ddk/strmini.h", "ddk/swenum.h", "ddk/tdikrnl.h", "ddk/tdistat.h", "ddk/upssvc.h", "ddk/usbbusif.h", "ddk/usbdlib.h", "ddk/usbdrivr.h", "ddk/usbkern.h", "ddk/usbprint.h", "ddk/usbprotocoldefs.h", "ddk/usbscan.h", "ddk/usbstorioctl.h", "ddk/video.h", "ddk/videoagp.h", "ddk/wdm.h", "ddk/wdmguid.h", "ddk/wmidata.h", "ddk/wmilib.h", "ddk/ws2san.h", "ddk/xfilter.h", "ddraw.h", "ddrawgdi.h", "ddrawi.h", "ddstream.h", "ddstream.idl", "debugapi.h", "delayimp.h", "devguid.h", "devicetopology.h", "devioctl.h", "devpkey.h", "devpropdef.h", "dhcpcsdk.h", "dhcpsapi.h", "dhcpssdk.h", "dhcpv6csdk.h", "dhtmldid.h", "dhtmled.h", "dhtmliid.h", "digitalv.h", "dimm.h", "dinput.h", "dir.h", "direct.h", "dirent.h", "diskguid.h", "dispatch.h", "dispdib.h", "dispex.h", "dlcapi.h", "dlgs.h", "dls1.h", "dls2.h", "dmdls.h", "dmemmgr.h", "dmerror.h", "dmksctrl.h", "dmo.h", "dmodshow.h", "dmodshow.idl", "dmoreg.h", "dmort.h", "dmplugin.h", "dmusbuff.h", "dmusicc.h", "dmusicf.h", "dmusici.h", "dmusics.h", "docobj.h", "docobjectservice.h", "documenttarget.h", "domdid.h", "dos.h", "downloadmgr.h", "dpaddr.h", "dpapi.h", "dpfilter.h", "dplay.h", "dplay8.h", "dplobby.h", "dplobby8.h", "dpnathlp.h", "driverspecs.h", "dsadmin.h", "dsclient.h", "dsconf.h", "dsdriver.h", "dsgetdc.h", "dshow.h", "dskquota.h", "dsound.h", "dsquery.h", "dsrole.h", "dssec.h", "dtchelp.h", "dvbsiparser.h", "dvdevcod.h", "dvdmedia.h", "dvec.h", "dvobj.h", "dwmapi.h", "dwrite.h", "dwrite_1.h", "dwrite_2.h", "dxdiag.h", "dxerr8.h", "dxerr9.h", "dxfile.h", "dxgi.h", "dxgi.idl", "dxgi1_2.h", "dxgi1_2.idl", "dxgiformat.h", "dxgitype.h", "dxtmpl.h", "dxva.h", "dxva2api.h", "dxvahd.h", "eapauthenticatoractiondefine.h", "eapauthenticatortypes.h", "eaphosterror.h", "eaphostpeerconfigapis.h", "eaphostpeertypes.h", "eapmethodauthenticatorapis.h", "eapmethodpeerapis.h", "eapmethodtypes.h", "eappapis.h", "eaptypes.h", "edevdefs.h", "eh.h", "ehstorapi.h", "elscore.h", "emostore.h", "emostore_i.c", "emptyvc.h", "endpointvolume.h", "errhandlingapi.h", "errno.h", "error.h", "errorrep.h", "errors.h", "esent.h", "evcode.h", "evcoll.h", "eventsys.h", "evntcons.h", "evntprov.h", "evntrace.h", "evr.h", "evr9.h", "exchform.h", "excpt.h", "exdisp.h", "exdispid.h", "fci.h", "fcntl.h", "fdi.h", "fenv.h", "fibersapi.h", "fileapi.h", "fileextd.h", "filehc.h", "fileopen.dlg", "filter.h", "filterr.h", "findtext.dlg", "float.h", "fltdefs.h", "fltuser.h", "fltuserstructures.h", "fltwinerror.h", "font.dlg", "fpieee.h", "fsrm.h", "fsrmenums.h", "fsrmerr.h", "fsrmpipeline.h", "fsrmquota.h", "fsrmreports.h", "fsrmscreen.h", "ftsiface.h", "ftw.h", "functiondiscoveryapi.h", "functiondiscoverycategories.h", "functiondiscoveryconstraints.h", "functiondiscoverykeys.h", "functiondiscoverykeys_devpkey.h", "functiondiscoverynotification.h", "fusion.h", "fvec.h", "fwpmtypes.h", "fwpmu.h", "fwptypes.h", "gb18030.h", "gdiplus.h", "gdiplus/gdiplus.h", "gdiplus/gdiplusbase.h", "gdiplus/gdiplusbrush.h", "gdiplus/gdipluscolor.h", "gdiplus/gdipluscolormatrix.h", "gdiplus/gdipluseffects.h", "gdiplus/gdiplusenums.h", "gdiplus/gdiplusflat.h", "gdiplus/gdiplusgpstubs.h", "gdiplus/gdiplusgraphics.h", "gdiplus/gdiplusheaders.h", "gdiplus/gdiplusimageattributes.h", "gdiplus/gdiplusimagecodec.h", "gdiplus/gdiplusimaging.h", "gdiplus/gdiplusimpl.h", "gdiplus/gdiplusinit.h", "gdiplus/gdipluslinecaps.h", "gdiplus/gdiplusmatrix.h", "gdiplus/gdiplusmem.h", "gdiplus/gdiplusmetafile.h", "gdiplus/gdiplusmetaheader.h", "gdiplus/gdipluspath.h", "gdiplus/gdipluspen.h", "gdiplus/gdipluspixelformats.h", "gdiplus/gdiplusstringformat.h", "gdiplus/gdiplustypes.h", "getopt.h", "gpedit.h", "gpio.h", "gpmgmt.h", "guiddef.h", "h323priv.h", "handleapi.h", "heapapi.h", "hidclass.h", "hidpi.h", "hidsdi.h", "hidusage.h", "highlevelmonitorconfigurationapi.h", "hlguids.h", "hliface.h", "hlink.h", "hostinfo.h", "hstring.h", "htiface.h", "htiframe.h", "htmlguid.h", "htmlhelp.h", "http.h", "httpext.h", "httpfilt.h", "httprequestid.h", "ia64reg.h", "iaccess.h", "iadmext.h", "iadmw.h", "iads.h", "icftypes.h", "icm.h", "icmpapi.h", "icmui.dlg", "icodecapi.h", "icrsint.h", "identitycommon.h", "identitystore.h", "idf.h", "idispids.h", "iedial.h", "ieeefp.h", "ieverp.h", "ifdef.h", "iiis.h", "iiisext.h", "iimgctx.h", "iiscnfg.h", "iisext_i.c", "iisrsta.h", "iketypes.h", "ilogobj.hxx", "imagehlp.h", "ime.h", "imessage.h", "imm.h", "in6addr.h", "inaddr.h", "indexsrv.h", "inetreg.h", "inetsdk.h", "infstr.h", "initguid.h", "initoid.h", "inputscope.h", "inspectable.h", "interlockedapi.h", "intrin.h", "intsafe.h", "intshcut.h", "inttypes.h", "invkprxy.h", "io.h", "ioapiset.h", "ioevent.h", "ipexport.h", "iphlpapi.h", "ipifcons.h", "ipinfoid.h", "ipmib.h", "ipmsp.h", "iprtrmib.h", "ipsectypes.h", "iptypes.h", "ipxconst.h", "ipxrip.h", "ipxrtdef.h", "ipxsap.h", "ipxtfflt.h", "iscsidsc.h", "isguids.h", "issper16.h", "issperr.h", "isysmon.h", "ivec.h", "iwamreg.h", "i_cryptasn1tls.h", "jobapi.h", "kcom.h", "knownfolders.h", "ks.h", "ksdebug.h", "ksguid.h", "ksmedia.h", "ksproxy.h", "ksuuids.h", "ktmtypes.h", "ktmw32.h", "kxia64.h", "l2cmn.h", "libgen.h", "libloaderapi.h", "limits.h", "lm.h", "lmaccess.h", "lmalert.h", "lmapibuf.h", "lmat.h", "lmaudit.h", "lmconfig.h", "lmcons.h", "lmdfs.h", "lmerr.h", "lmerrlog.h", "lmjoin.h", "lmmsg.h", "lmon.h", "lmremutl.h", "lmrepl.h", "lmserver.h", "lmshare.h", "lmsname.h", "lmstats.h", "lmsvc.h", "lmuse.h", "lmuseflg.h", "lmwksta.h", "loadperf.h", "locale.h", "locationapi.h", "lpmapi.h", "lzexpand.h", "madcapcl.h", "magnification.h", "mailmsgprops.h", "malloc.h", "manipulations.h", "mapi.h", "mapicode.h", "mapidbg.h", "mapidefs.h", "mapiform.h", "mapiguid.h", "mapihook.h", "mapinls.h", "mapioid.h", "mapispi.h", "mapitags.h", "mapiutil.h", "mapival.h", "mapiwin.h", "mapiwz.h", "mapix.h", "math.h", "mbctype.h", "mbstring.h", "mciavi.h", "mcx.h", "mdbrole.hxx", "mdcommsg.h", "mddefw.h", "mdhcp.h", "mdmsg.h", "mediaerr.h", "mediaobj.h", "mediaobj.idl", "medparam.h", "medparam.idl", "mem.h", "memory.h", "memoryapi.h", "mergemod.h", "mfapi.h", "mferror.h", "mfidl.h", "mfmp2dlna.h", "mfobjects.h", "mfplay.h", "mfreadwrite.h", "mftransform.h", "mgm.h", "mgmtapi.h", "midles.h", "mimedisp.h", "mimeinfo.h", "minmax.h", "minwinbase.h", "minwindef.h", "mlang.h", "mmc.h", "mmcobj.h", "mmdeviceapi.h", "mmreg.h", "mmstream.h", "mmstream.idl", "mmsystem.h", "mobsync.h", "moniker.h", "mpeg2bits.h", "mpeg2data.h", "mpeg2psiparser.h", "mpeg2structs.h", "mprapi.h", "mprerror.h", "mq.h", "mqmail.h", "mqoai.h", "msacm.h", "msacmdlg.dlg", "msacmdlg.h", "msado15.h", "msasn1.h", "msber.h", "mscat.h", "mschapp.h", "msclus.h", "mscoree.h", "msctf.h", "msctfmonitorapi.h", "msdadc.h", "msdaguid.h", "msdaipp.h", "msdaipper.h", "msdaora.h", "msdaosp.h", "msdasc.h", "msdasql.h", "msdatsrc.h", "msdrm.h", "msdrmdefs.h", "msdshape.h", "msfs.h", "mshtmcid.h", "mshtmdid.h", "mshtmhst.h", "mshtml.h", "mshtmlc.h", "msi.h", "msidefs.h", "msimcntl.h", "msimcsdk.h", "msinkaut.h", "msinkaut_i.c", "msiquery.h", "msoav.h", "msopc.h", "msp.h", "mspab.h", "mspaddr.h", "mspbase.h", "mspcall.h", "mspcoll.h", "mspenum.h", "msplog.h", "mspst.h", "mspstrm.h", "mspterm.h", "mspthrd.h", "msptrmac.h", "msptrmar.h", "msptrmvc.h", "msputils.h", "msrdc.h", "msremote.h", "mssip.h", "msstkppg.h", "mstask.h", "mstcpip.h", "msterr.h", "mswsock.h", "msxml.h", "msxml2.h", "msxml2did.h", "msxmldid.h", "mtsadmin.h", "mtsadmin_i.c", "mtsevents.h", "mtsgrp.h", "mtx.h", "mtxadmin.h", "mtxadmin_i.c", "mtxattr.h", "mtxdm.h", "muiload.h", "multimon.h", "multinfo.h", "mxdc.h", "namedpipeapi.h", "namespaceapi.h", "napcertrelyingparty.h", "napcommon.h", "napenforcementclient.h", "napmanagement.h", "napmicrosoftvendorids.h", "napprotocol.h", "napservermanagement.h", "napsystemhealthagent.h", "napsystemhealthvalidator.h", "naptypes.h", "naputil.h", "nb30.h", "ncrypt.h", "ndattrib.h", "ndfapi.h", "ndhelper.h", "ndkinfo.h", "ndr64types.h", "ndrtypes.h", "netcon.h", "neterr.h", "netevent.h", "netioapi.h", "netlistmgr.h", "netmon.h", "netprov.h", "nettypes.h", "new.h", "newapis.h", "newdev.h", "nldef.h", "nmsupp.h", "npapi.h", "nsemail.h", "nspapi.h", "ntdd1394.h", "ntdd8042.h", "ntddbeep.h", "ntddcdrm.h", "ntddcdvd.h", "ntddchgr.h", "ntdddisk.h", "ntddft.h", "ntddkbd.h", "ntddmmc.h", "ntddmodm.h", "ntddmou.h", "ntddndis.h", "ntddpar.h", "ntddpsch.h", "ntddscsi.h", "ntddser.h", "ntddstor.h", "ntddtape.h", "ntddtdi.h", "ntddvdeo.h", "ntddvol.h", "ntdef.h", "ntdsapi.h", "ntdsbcli.h", "ntdsbmsg.h", "ntgdi.h", "ntiologc.h", "ntldap.h", "ntmsapi.h", "ntmsmli.h", "ntquery.h", "ntsdexts.h", "ntsecapi.h", "ntsecpkg.h", "ntstatus.h", "ntverp.h", "oaidl.h", "objbase.h", "objectarray.h", "objerror.h", "objidl.h", "objidlbase.h", "objsafe.h", "objsel.h", "ocidl.h", "ocmm.h", "odbcinst.h", "odbcss.h", "ole.h", "ole2.h", "ole2ver.h", "oleacc.h", "oleauto.h", "olectl.h", "olectlid.h", "oledb.h", "oledbdep.h", "oledberr.h", "oledbguid.h", "oledlg.dlg", "oledlg.h", "oleidl.h", "oletx2xa.h", "opmapi.h", "optary.h", "p2p.h", "packoff.h", "packon.h", "parser.h", "patchapi.h", "patchwiz.h", "pathcch.h", "pbt.h", "pchannel.h", "pciprop.h", "pcrt32.h", "pdh.h", "pdhmsg.h", "penwin.h", "perflib.h", "perhist.h", "persist.h", "pgobootrun.h", "physicalmonitorenumerationapi.h", "pla.h", "pnrpdef.h", "pnrpns.h", "poclass.h", "polarity.h", "poppack.h", "portabledeviceconnectapi.h", "portabledevicetypes.h", "powrprof.h", "prnasnot.h", "prnsetup.dlg", "prntfont.h", "process.h", "processenv.h", "processthreadsapi.h", "processtopologyapi.h", "profile.h", "profileapi.h", "profinfo.h", "propidl.h", "propkey.h", "propkeydef.h", "propsys.h", "propvarutil.h", "prsht.h", "psapi.h", "psdk_inc/intrin-impl.h", "psdk_inc/_dbg_LOAD_IMAGE.h", "psdk_inc/_dbg_common.h", "psdk_inc/_fd_types.h", "psdk_inc/_ip_mreq1.h", "psdk_inc/_ip_types.h", "psdk_inc/_pop_BOOL.h", "psdk_inc/_push_BOOL.h", "psdk_inc/_socket_types.h", "psdk_inc/_varenum.h", "psdk_inc/_ws1_undef.h", "psdk_inc/_wsadata.h", "psdk_inc/_wsa_errnos.h", "psdk_inc/_xmitfile.h", "pshpack1.h", "pshpack2.h", "pshpack4.h", "pshpack8.h", "pshpck16.h", "pstore.h", "pthread.h", "pthread_compat.h", "pthread_signal.h", "pthread_time.h", "pthread_unistd.h", "qedit.h", "qedit.idl", "qmgr.h", "qnetwork.h", "qnetwork.idl", "qos.h", "qos2.h", "qosname.h", "qospol.h", "qossp.h", "ras.h", "rasdlg.h", "raseapif.h", "raserror.h", "rassapi.h", "rasshost.h", "ratings.h", "rdpencomapi.h", "realtimeapiset.h", "reason.h", "recguids.h", "reconcil.h", "regbag.h", "regstr.h", "rend.h", "resapi.h", "restartmanager.h", "richedit.h", "richole.h", "rkeysvcc.h", "rnderr.h", "roapi.h", "routprot.h", "rpc.h", "rpcasync.h", "rpcdce.h", "rpcdcep.h", "rpcndr.h", "rpcnsi.h", "rpcnsip.h", "rpcnterr.h", "rpcproxy.h", "rpcsal.h", "rpcssl.h", "rrascfg.h", "rtcapi.h", "rtccore.h", "rtcerr.h", "rtinfo.h", "rtm.h", "rtmv2.h", "rtutils.h", "sal.h", "sapi.h", "sapi51.h", "sapi53.h", "sapi54.h", "sas.h", "sbe.h", "scarddat.h", "scarderr.h", "scardmgr.h", "scardsrv.h", "scardssp.h", "scardssp_i.c", "scardssp_p.c", "scesvc.h", "schannel.h", "sched.h", "schedule.h", "schemadef.h", "schnlsp.h", "scode.h", "scrnsave.h", "scrptids.h", "sddl.h", "sdkddkver.h", "sdks/_mingw_ddk.h", "sdks/_mingw_directx.h", "sdoias.h", "sdpblb.h", "sdperr.h", "search.h", "secext.h", "security.h", "securityappcontainer.h", "securitybaseapi.h", "sec_api/conio_s.h", "sec_api/crtdbg_s.h", "sec_api/mbstring_s.h", "sec_api/search_s.h", "sec_api/stdio_s.h", "sec_api/stdlib_s.h", "sec_api/stralign_s.h", "sec_api/string_s.h", "sec_api/sys/timeb_s.h", "sec_api/tchar_s.h", "sec_api/wchar_s.h", "sehmap.h", "semaphore.h", "sens.h", "sensapi.h", "sensevts.h", "sensors.h", "sensorsapi.h", "servprov.h", "setjmp.h", "setjmpex.h", "setupapi.h", "sfc.h", "shappmgr.h", "share.h", "shdeprecated.h", "shdispid.h", "shellapi.h", "sherrors.h", "shfolder.h", "shldisp.h", "shlguid.h", "shlobj.h", "shlwapi.h", "shobjidl.h", "shtypes.h", "signal.h", "simpdata.h", "simpdc.h", "sipbase.h", "sisbkup.h", "slerror.h", "slpublic.h", "smpab.h", "smpms.h", "smpxp.h", "smtpguid.h", "smx.h", "snmp.h", "softpub.h", "specstrings.h", "sperror.h", "sphelper.h", "sporder.h", "sql.h", "sqlext.h", "sqloledb.h", "sqltypes.h", "sqlucode.h", "sql_1.h", "srrestoreptapi.h", "srv.h", "sspguid.h", "sspi.h", "sspserr.h", "sspsidl.h", "stdarg.h", "stddef.h", "stdexcpt.h", "stdint.h", "stdio.h", "stdlib.h", "sti.h", "stierr.h", "stireg.h", "stllock.h", "stm.h", "storage.h", "storduid.h", "storprop.h", "stralign.h", "string.h", "stringapiset.h", "strings.h", "strmif.h", "strsafe.h", "structuredquerycondition.h", "subauth.h", "subsmgr.h", "svcguid.h", "svrapi.h", "swprintf.inl", "synchapi.h", "sysinfoapi.h", "syslimits.h", "systemtopologyapi.h", "sys/cdefs.h", "sys/fcntl.h", "sys/file.h", "sys/locking.h", "sys/param.h", "sys/stat.h", "sys/time.h", "sys/timeb.h", "sys/types.h", "sys/unistd.h", "sys/utime.h", "t2embapi.h", "tabflicks.h", "tapi.h", "tapi3.h", "tapi3cc.h", "tapi3ds.h", "tapi3err.h", "tapi3if.h", "taskschd.h", "tbs.h", "tcerror.h", "tcguid.h", "tchar.h", "tcpestats.h", "tcpmib.h", "tdh.h", "tdi.h", "tdiinfo.h", "termmgr.h", "textserv.h", "textstor.h", "threadpoolapiset.h", "threadpoollegacyapiset.h", "time.h", "timeprov.h", "timezoneapi.h", "tlbref.h", "tlhelp32.h", "tlogstg.h", "tmschema.h", "tnef.h", "tom.h", "tpcshrd.h", "traffic.h", "transact.h", "triedcid.h", "triediid.h", "triedit.h", "tsattrs.h", "tspi.h", "tssbx.h", "tsuserex.h", "tsuserex_i.c", "tuner.h", "tvout.h", "txcoord.h", "txctx.h", "txdtc.h", "txfw32.h", "typeinfo.h", "uastrfnc.h", "uchar.h", "udpmib.h", "uiautomation.h", "uiautomationclient.h", "uiautomationcore.h", "uiautomationcoreapi.h", "uiviewsettingsinterop.h", "umx.h", "unistd.h", "unknown.h", "unknwn.h", "unknwnbase.h", "urlhist.h", "urlmon.h", "usb.h", "usb100.h", "usb200.h", "usbcamdi.h", "usbdi.h", "usbioctl.h", "usbiodef.h", "usbprint.h", "usbrpmif.h", "usbscan.h", "usbspec.h", "usbuser.h", "userenv.h", "usp10.h", "utilapiset.h", "utime.h", "uuids.h", "uxtheme.h", "vadefs.h", "varargs.h", "vcr.h", "vdmdbg.h", "vds.h", "vdslun.h", "verinfo.ver", "versionhelpers.h", "vfw.h", "vfwmsgs.h", "virtdisk.h", "vmr9.h", "vmr9.idl", "vsadmin.h", "vsbackup.h", "vsmgmt.h", "vsprov.h", "vss.h", "vsstyle.h", "vssym32.h", "vswriter.h", "w32api.h", "wab.h", "wabapi.h", "wabcode.h", "wabdefs.h", "wabiab.h", "wabmem.h", "wabnot.h", "wabtags.h", "wabutil.h", "wbemads.h", "wbemcli.h", "wbemdisp.h", "wbemidl.h", "wbemprov.h", "wbemtran.h", "wchar.h", "wcmconfig.h", "wcsplugin.h", "wct.h", "wctype.h", "wdsbp.h", "wdsclientapi.h", "wdspxe.h", "wdstci.h", "wdstpdi.h", "wdstptmgmt.h", "werapi.h", "wfext.h", "wia.h", "wiadef.h", "wiadevd.h", "wiavideo.h", "winable.h", "winapifamily.h", "winbase.h", "winber.h", "wincodec.h", "wincon.h", "wincred.h", "wincrypt.h", "winddi.h", "winddiui.h", "windef.h", "windns.h", "windot11.h", "windows.foundation.h", "windows.h", "windows.security.cryptography.h", "windows.storage.h", "windows.storage.streams.h", "windows.system.threading.h", "windowsx.h", "windowsx.h16", "winefs.h", "winerror.h", "winevt.h", "wingdi.h", "winhttp.h", "wininet.h", "winineti.h", "winioctl.h", "winldap.h", "winnetwk.h", "winnls.h", "winnls32.h", "winnt.h", "winnt.rh", "winperf.h", "winreg.h", "winresrc.h", "winsafer.h", "winsatcominterfacei.h", "winscard.h", "winsdkver.h", "winsmcrd.h", "winsnmp.h", "winsock.h", "winsock2.h", "winsplp.h", "winspool.h", "winstring.h", "winsvc.h", "winsxs.h", "winsync.h", "winternl.h", "wintrust.h", "winusb.h", "winusbio.h", "winuser.h", "winuser.rh", "winver.h", "winwlx.h", "wlanapi.h", "wlanihvtypes.h", "wlantypes.h", "wmcodecdsp.h", "wmcontainer.h", "wmiatlprov.h", "wmistr.h", "wmiutils.h", "wmsbuffer.h", "wmsdkidl.h", "wnnc.h", "wow64apiset.h", "wownt16.h", "wownt32.h", "wpapi.h", "wpapimsg.h", "wpcapi.h", "wpcevent.h", "wpcrsmsg.h", "wpftpmsg.h", "wppstmsg.h", "wpspihlp.h", "wptypes.h", "wpwizmsg.h", "wrl.h", "wrl/client.h", "wrl/internal.h", "wrl/module.h", "wrl/wrappers/corewrappers.h", "ws2atm.h", "ws2bth.h", "ws2def.h", "ws2dnet.h", "ws2ipdef.h", "ws2spi.h", "ws2tcpip.h", "wsdapi.h", "wsdattachment.h", "wsdbase.h", "wsdclient.h", "wsddisco.h", "wsdhost.h", "wsdtypes.h", "wsdutil.h", "wsdxml.h", "wsdxmldom.h", "wshisotp.h", "wsipv6ok.h", "wsipx.h", "wsman.h", "wsmandisp.h", "wsnetbs.h", "wsnwlink.h", "wspiapi.h", "wsrm.h", "wsvns.h", "wtsapi32.h", "wtypes.h", "wtypesbase.h", "xa.h", "xcmc.h", "xcmcext.h", "xcmcmsx2.h", "xcmcmsxt.h", "xenroll.h", "xinput.h", "xlocinfo.h", "xmath.h", "xmldomdid.h", "xmldsodid.h", "xmllite.h", "xmltrnsf.h", "xolehlp.h", "xpsdigitalsignature.h", "xpsobjectmodel.h", "xpsobjectmodel_1.h", "xpsprint.h", "xpsrassvc.h", "ymath.h", "yvals.h", "zmouse.h", "_bsd_types.h", "_cygwin.h", "_dbdao.h", "_mingw.h", "_mingw_dxhelper.h", "_mingw_mac.h", "_mingw_off_t.h", "_mingw_print_pop.h", "_mingw_print_push.h", "_mingw_secapi.h", "_mingw_stat64.h", "_mingw_stdarg.h", "_mingw_unicode.h", "_timeval.h"}...)
	}
	return xs
}

func NewLocalSystem(verbose bool) (*LocalSystem, error) {
	var locsys LocalSystem
	locsys.includeFiles = make([]string, 0)
	locsys.systemIncludeDirectories = locsys.SystemIncludeDirectories()
	locsys.commonIncludes = locsys.CommonIncludes()
	locsys.localIncludeDirectories = []string{".", "include", "Include", "..", "../include", "../Include", "common", "Common", "../common", "../Common"}
	for _, rootPath := range locsys.systemIncludeDirectories {
		err := filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			switch strings.ToLower(filepath.Ext(path)) {
			case ".h", ".hpp", ".hh", ".h++":
				//if verbose {
				//fmt.Printf("added: %q\n", path)
				//}

				locsys.includeFiles = append(locsys.includeFiles, path)
			}
			for _, commonInclude := range locsys.commonIncludes {
				if strings.HasSuffix(path, commonInclude) {
					//if verbose {
					//fmt.Printf("added: %q\n", path)
					//}
					locsys.includeFiles = append(locsys.includeFiles, path)
				}
			}
			return nil
		})
		if err != nil {
			if verbose {
				fmt.Println(err)
			}
		}
	}
	if verbose {
		fmt.Printf("Found %d include files in %s\n", len(locsys.includeFiles), strings.Join(locsys.systemIncludeDirectories, ", "))
	}
	locsys.verbose = verbose
	return &locsys, nil
}

func (locsys *LocalSystem) IncludeFiles() []string {
	return locsys.includeFiles
}
