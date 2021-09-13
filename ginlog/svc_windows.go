/**
 * Auth :   liubo
 * Date :   2021/9/13 13:42
 * Comment:
 */

package ginlog

import "golang.org/x/sys/windows/svc"

func init() {
	svcIsWindowsService = svc.IsWindowsService
}
