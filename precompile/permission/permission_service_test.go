package permission

import (
	"testing"
	"crypto/ecdsa"

	"github.com/KasperLiu/gobcos/client"
	"github.com/KasperLiu/gobcos/crypto"
)

const (
	success = "{\"code\":0,\"msg\":\"success\"}"
	tableName = "t_test"
	permisstionAdd = "0xFbb18d54e9Ee57529cda8c7c52242EFE879f064F"
	txOrigin = "0x06527eA53361EE68D4F671D4b50FB6B2D82Dcb23"
)

func GetClient(t *testing.T) *client.Client {
	groupID := uint(1)
	rpc, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		t.Fatalf("init rpc client failed: %+v", err)
	}
	return rpc
}

func GenerateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("608fe45cc95cce1b5b048ea588cfab5936fd5ed7cdb19dfe68404d1a462ef5ab")
    if err != nil {
        t.Fatalf("init privateKey failed: %+v", err)
	}
	return privateKey
}

func GetService(t *testing.T) *PermissionService {
	rpc := GetClient(t)
	privateKey := GenerateKey(t)
	service, err := NewPermissionService(rpc, privateKey)
	if err != nil {
		t.Fatalf("init PermissionService failed: %+v", err)
	}
	return service
}

func TestGrant(t *testing.T) {
	service := GetService(t)
	// grant permission
	// result, err := service.GrantPermissionManager(permisstionAdd)
	// if err != nil {
	// 	t.Fatalf("TestPermissionManager failed: %v", err)
	// }
	// t.Logf("TestPermissionManager: %v", result)
	result, err := service.RevokePermissionManager(txOrigin)
	if err != nil {
		t.Fatalf("RevokePermissionManager failed: %v", err)
	}
	t.Logf("RevokePermissionManager: %v", result)

	// result, err := service.ListPermissionManager()
	// if err != nil {
	// 	t.Fatalf("ListPermissionManager failed: %v", err)
	// }
	// t.Logf("ListPermissionManager: %v", result)
}

// func TestPermissionManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantPermissionManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestPermissionManager failed: %v", err)
// 	}
// 	t.Logf("TestPermissionManager: %v", result)
// 	revokeResult, err := service.RevokePermissionManager(permisstionAdd)
// 	if err != nil {
// 		t.Fatalf("TestPermissionManager failed: %v", err)
// 	}
// 	t.Logf("TestPermissionManager revoke result: %v", revokeResult)
// }

// func TestUserTableManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantUserTableManager(tableName, txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestUserTableManager failed: %v", err)
// 	}
// 	t.Logf("TestUserTableManager: %v", result)
// 	revokeResult, err := service.RevokeUserTableManager(tableName, txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestUserTableManager failed: %v", err)
// 	}
// 	t.Logf("TestUserTableManager revoke result: %v", revokeResult)
// }

// func TestDeployAndCreateManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantDeployAndCreateManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestDeployAndCreateManager failed: %v", err)
// 	}
// 	t.Logf("TestDeployAndCreateManager: %v", result)

// 	revokeResult, err := service.RevokeDeployAndCreateManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestDeployAndCreateManager failed: %v", err)
// 	}
// 	t.Logf("TestDeployAndCreateManager revoke result: %v", revokeResult)
// }

// func TestNodeManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantNodeManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestNodeManager failed: %v", err)
// 	}
// 	t.Logf("TestNodeManager: %v", result)

// 	revokeResult, err := service.RevokeNodeManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestNodeManager failed: %v", err)
// 	}
// 	t.Logf("TestNodeManager revoke result: %v", revokeResult)
// }

// func TestCNSManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantCNSManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestCNSManager failed: %v", err)
// 	}
// 	t.Logf("TestCNSManager: %v", result)

// 	revokeResult, err := service.RevokeCNSManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestCNSManager failed: %v", err)
// 	}
// 	t.Logf("TestCNSManager revoke result: %v", revokeResult)
// }

// func TestSysConfigManager(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.GrantSysConfigManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestSysConfigManager failed: %v", err)
// 	}
// 	t.Logf("TestSysConfigManager: %v", result)

// 	revokeResult, err := service.RevokeSysConfigManager(txOrigin)
// 	if err != nil {
// 		t.Fatalf("TestSysConfigManager failed: %v", err)
// 	}
// 	t.Logf("TestSysConfigManager revoke result: %v", revokeResult)
// 	t.Logf("Success result: %s", success)
// }

// func TestListUser(t *testing.T) {
// 	service := GetService(t)

// 	result, err := service.ListUserTableManager(tableName)
// 	if err != nil {
// 		t.Fatalf("ListUserTableManager failed: %v", err)
// 	}
// 	t.Logf("ListUserTableManager: %v", result)
// }