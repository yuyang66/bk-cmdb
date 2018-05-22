package object

import (
	"configcenter/src/common/errors"
	"configcenter/src/scene_server/topo_server/topo_service/manager"
	api "configcenter/src/source_controller/api/object"
)

type topoGraphicsLogic struct {
	objcli *api.Client
	cfg    manager.Configer
	mgr    manager.Manager
}

var _ manager.TopoGraphics = (*topoGraphicsLogic)(nil) // check the interface

func init() {
	obj := &topoGraphicsLogic{}

	obj.objcli = api.NewClient("")
	manager.SetManager(obj)
	manager.RegisterLogic(manager.ObjectAsst, obj)
}

// Set implement SetConfiger interface
func (cli *topoGraphicsLogic) Set(cfg manager.Configer) {
	cli.cfg = cfg
}

// SetManager implement the manager's Hooker interface
func (cli *topoGraphicsLogic) SetManager(mgr manager.Manager) error {
	cli.mgr = mgr
	return nil
}

func (cli *topoGraphicsLogic) SearchGraphics(forward *api.ForwardParam, params map[string]interface{}, errProxy errors.DefaultCCErrorIf) ([]api.TopoGraphics, error) {
	cli.objcli.SetAddress(cli.cfg.Get(cli))
	return cli.objcli.SearchTopoGraphics(forward, params)
}
