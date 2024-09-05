// Code generated from DmSqlParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DmSqlParser
import "github.com/antlr4-go/antlr/v4"

// BaseDmSqlParserListener is a complete listener for a parse tree produced by DmSqlParser.
type BaseDmSqlParserListener struct{}

var _ DmSqlParserListener = &BaseDmSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDmSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDmSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDmSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDmSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDmprogram is called when production dmprogram is entered.
func (s *BaseDmSqlParserListener) EnterDmprogram(ctx *DmprogramContext) {}

// ExitDmprogram is called when production dmprogram is exited.
func (s *BaseDmSqlParserListener) ExitDmprogram(ctx *DmprogramContext) {}

// EnterSql_clauses is called when production sql_clauses is entered.
func (s *BaseDmSqlParserListener) EnterSql_clauses(ctx *Sql_clausesContext) {}

// ExitSql_clauses is called when production sql_clauses is exited.
func (s *BaseDmSqlParserListener) ExitSql_clauses(ctx *Sql_clausesContext) {}

// EnterDdlsql is called when production ddlsql is entered.
func (s *BaseDmSqlParserListener) EnterDdlsql(ctx *DdlsqlContext) {}

// ExitDdlsql is called when production ddlsql is exited.
func (s *BaseDmSqlParserListener) ExitDdlsql(ctx *DdlsqlContext) {}

// EnterDmlsql is called when production dmlsql is entered.
func (s *BaseDmSqlParserListener) EnterDmlsql(ctx *DmlsqlContext) {}

// ExitDmlsql is called when production dmlsql is exited.
func (s *BaseDmSqlParserListener) ExitDmlsql(ctx *DmlsqlContext) {}

// EnterPrivsql is called when production privsql is entered.
func (s *BaseDmSqlParserListener) EnterPrivsql(ctx *PrivsqlContext) {}

// ExitPrivsql is called when production privsql is exited.
func (s *BaseDmSqlParserListener) ExitPrivsql(ctx *PrivsqlContext) {}

// EnterOthersql is called when production othersql is entered.
func (s *BaseDmSqlParserListener) EnterOthersql(ctx *OthersqlContext) {}

// ExitOthersql is called when production othersql is exited.
func (s *BaseDmSqlParserListener) ExitOthersql(ctx *OthersqlContext) {}

// EnterUtilsql is called when production utilsql is entered.
func (s *BaseDmSqlParserListener) EnterUtilsql(ctx *UtilsqlContext) {}

// ExitUtilsql is called when production utilsql is exited.
func (s *BaseDmSqlParserListener) ExitUtilsql(ctx *UtilsqlContext) {}

// EnterExplainsql is called when production explainsql is entered.
func (s *BaseDmSqlParserListener) EnterExplainsql(ctx *ExplainsqlContext) {}

// ExitExplainsql is called when production explainsql is exited.
func (s *BaseDmSqlParserListener) ExitExplainsql(ctx *ExplainsqlContext) {}

// EnterShutdown_stmt is called when production shutdown_stmt is entered.
func (s *BaseDmSqlParserListener) EnterShutdown_stmt(ctx *Shutdown_stmtContext) {}

// ExitShutdown_stmt is called when production shutdown_stmt is exited.
func (s *BaseDmSqlParserListener) ExitShutdown_stmt(ctx *Shutdown_stmtContext) {}

// EnterAlter_diskgroup_stmt is called when production alter_diskgroup_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_diskgroup_stmt(ctx *Alter_diskgroup_stmtContext) {}

// ExitAlter_diskgroup_stmt is called when production alter_diskgroup_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_diskgroup_stmt(ctx *Alter_diskgroup_stmtContext) {}

// EnterLocal is called when production local is entered.
func (s *BaseDmSqlParserListener) EnterLocal(ctx *LocalContext) {}

// ExitLocal is called when production local is exited.
func (s *BaseDmSqlParserListener) ExitLocal(ctx *LocalContext) {}

// EnterDmsubprogram is called when production dmsubprogram is entered.
func (s *BaseDmSqlParserListener) EnterDmsubprogram(ctx *DmsubprogramContext) {}

// ExitDmsubprogram is called when production dmsubprogram is exited.
func (s *BaseDmSqlParserListener) ExitDmsubprogram(ctx *DmsubprogramContext) {}

// EnterDeclare_block is called when production declare_block is entered.
func (s *BaseDmSqlParserListener) EnterDeclare_block(ctx *Declare_blockContext) {}

// ExitDeclare_block is called when production declare_block is exited.
func (s *BaseDmSqlParserListener) ExitDeclare_block(ctx *Declare_blockContext) {}

// EnterDecl_var_cur_list_options is called when production decl_var_cur_list_options is entered.
func (s *BaseDmSqlParserListener) EnterDecl_var_cur_list_options(ctx *Decl_var_cur_list_optionsContext) {
}

// ExitDecl_var_cur_list_options is called when production decl_var_cur_list_options is exited.
func (s *BaseDmSqlParserListener) ExitDecl_var_cur_list_options(ctx *Decl_var_cur_list_optionsContext) {
}

// EnterDecl_var_cur_list_1 is called when production decl_var_cur_list_1 is entered.
func (s *BaseDmSqlParserListener) EnterDecl_var_cur_list_1(ctx *Decl_var_cur_list_1Context) {}

// ExitDecl_var_cur_list_1 is called when production decl_var_cur_list_1 is exited.
func (s *BaseDmSqlParserListener) ExitDecl_var_cur_list_1(ctx *Decl_var_cur_list_1Context) {}

// EnterDecl_var_cur_list is called when production decl_var_cur_list is entered.
func (s *BaseDmSqlParserListener) EnterDecl_var_cur_list(ctx *Decl_var_cur_listContext) {}

// ExitDecl_var_cur_list is called when production decl_var_cur_list is exited.
func (s *BaseDmSqlParserListener) ExitDecl_var_cur_list(ctx *Decl_var_cur_listContext) {}

// EnterDecl_plsql_type is called when production decl_plsql_type is entered.
func (s *BaseDmSqlParserListener) EnterDecl_plsql_type(ctx *Decl_plsql_typeContext) {}

// ExitDecl_plsql_type is called when production decl_plsql_type is exited.
func (s *BaseDmSqlParserListener) ExitDecl_plsql_type(ctx *Decl_plsql_typeContext) {}

// EnterPlsql_type_def is called when production plsql_type_def is entered.
func (s *BaseDmSqlParserListener) EnterPlsql_type_def(ctx *Plsql_type_defContext) {}

// ExitPlsql_type_def is called when production plsql_type_def is exited.
func (s *BaseDmSqlParserListener) ExitPlsql_type_def(ctx *Plsql_type_defContext) {}

// EnterLt_int_lst is called when production lt_int_lst is entered.
func (s *BaseDmSqlParserListener) EnterLt_int_lst(ctx *Lt_int_lstContext) {}

// ExitLt_int_lst is called when production lt_int_lst is exited.
func (s *BaseDmSqlParserListener) ExitLt_int_lst(ctx *Lt_int_lstContext) {}

// EnterRec_item_def_list is called when production rec_item_def_list is entered.
func (s *BaseDmSqlParserListener) EnterRec_item_def_list(ctx *Rec_item_def_listContext) {}

// ExitRec_item_def_list is called when production rec_item_def_list is exited.
func (s *BaseDmSqlParserListener) ExitRec_item_def_list(ctx *Rec_item_def_listContext) {}

// EnterRec_item_def is called when production rec_item_def is entered.
func (s *BaseDmSqlParserListener) EnterRec_item_def(ctx *Rec_item_defContext) {}

// ExitRec_item_def is called when production rec_item_def is exited.
func (s *BaseDmSqlParserListener) ExitRec_item_def(ctx *Rec_item_defContext) {}

// EnterDecl_variable is called when production decl_variable is entered.
func (s *BaseDmSqlParserListener) EnterDecl_variable(ctx *Decl_variableContext) {}

// ExitDecl_variable is called when production decl_variable is exited.
func (s *BaseDmSqlParserListener) ExitDecl_variable(ctx *Decl_variableContext) {}

// EnterNot_null is called when production not_null is entered.
func (s *BaseDmSqlParserListener) EnterNot_null(ctx *Not_nullContext) {}

// ExitNot_null is called when production not_null is exited.
func (s *BaseDmSqlParserListener) ExitNot_null(ctx *Not_nullContext) {}

// EnterPlsql_datatype is called when production plsql_datatype is entered.
func (s *BaseDmSqlParserListener) EnterPlsql_datatype(ctx *Plsql_datatypeContext) {}

// ExitPlsql_datatype is called when production plsql_datatype is exited.
func (s *BaseDmSqlParserListener) ExitPlsql_datatype(ctx *Plsql_datatypeContext) {}

// EnterDefault_clause_option is called when production default_clause_option is entered.
func (s *BaseDmSqlParserListener) EnterDefault_clause_option(ctx *Default_clause_optionContext) {}

// ExitDefault_clause_option is called when production default_clause_option is exited.
func (s *BaseDmSqlParserListener) ExitDefault_clause_option(ctx *Default_clause_optionContext) {}

// EnterVariable_name_list is called when production variable_name_list is entered.
func (s *BaseDmSqlParserListener) EnterVariable_name_list(ctx *Variable_name_listContext) {}

// ExitVariable_name_list is called when production variable_name_list is exited.
func (s *BaseDmSqlParserListener) ExitVariable_name_list(ctx *Variable_name_listContext) {}

// EnterDecl_except is called when production decl_except is entered.
func (s *BaseDmSqlParserListener) EnterDecl_except(ctx *Decl_exceptContext) {}

// ExitDecl_except is called when production decl_except is exited.
func (s *BaseDmSqlParserListener) ExitDecl_except(ctx *Decl_exceptContext) {}

// EnterPragma_def is called when production pragma_def is entered.
func (s *BaseDmSqlParserListener) EnterPragma_def(ctx *Pragma_defContext) {}

// ExitPragma_def is called when production pragma_def is exited.
func (s *BaseDmSqlParserListener) ExitPragma_def(ctx *Pragma_defContext) {}

// EnterPragma is called when production pragma is entered.
func (s *BaseDmSqlParserListener) EnterPragma(ctx *PragmaContext) {}

// ExitPragma is called when production pragma is exited.
func (s *BaseDmSqlParserListener) ExitPragma(ctx *PragmaContext) {}

// EnterPlbody is called when production plbody is entered.
func (s *BaseDmSqlParserListener) EnterPlbody(ctx *PlbodyContext) {}

// ExitPlbody is called when production plbody is exited.
func (s *BaseDmSqlParserListener) ExitPlbody(ctx *PlbodyContext) {}

// EnterSs_plbody is called when production ss_plbody is entered.
func (s *BaseDmSqlParserListener) EnterSs_plbody(ctx *Ss_plbodyContext) {}

// ExitSs_plbody is called when production ss_plbody is exited.
func (s *BaseDmSqlParserListener) ExitSs_plbody(ctx *Ss_plbodyContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseDmSqlParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseDmSqlParserListener) ExitLabel(ctx *LabelContext) {}

// EnterLabel_list is called when production label_list is entered.
func (s *BaseDmSqlParserListener) EnterLabel_list(ctx *Label_listContext) {}

// ExitLabel_list is called when production label_list is exited.
func (s *BaseDmSqlParserListener) ExitLabel_list(ctx *Label_listContext) {}

// EnterLabel_list_options is called when production label_list_options is entered.
func (s *BaseDmSqlParserListener) EnterLabel_list_options(ctx *Label_list_optionsContext) {}

// ExitLabel_list_options is called when production label_list_options is exited.
func (s *BaseDmSqlParserListener) ExitLabel_list_options(ctx *Label_list_optionsContext) {}

// EnterLabel_demiliter_l is called when production label_demiliter_l is entered.
func (s *BaseDmSqlParserListener) EnterLabel_demiliter_l(ctx *Label_demiliter_lContext) {}

// ExitLabel_demiliter_l is called when production label_demiliter_l is exited.
func (s *BaseDmSqlParserListener) ExitLabel_demiliter_l(ctx *Label_demiliter_lContext) {}

// EnterLabel_demiliter_r is called when production label_demiliter_r is entered.
func (s *BaseDmSqlParserListener) EnterLabel_demiliter_r(ctx *Label_demiliter_rContext) {}

// ExitLabel_demiliter_r is called when production label_demiliter_r is exited.
func (s *BaseDmSqlParserListener) ExitLabel_demiliter_r(ctx *Label_demiliter_rContext) {}

// EnterPlsql is called when production plsql is entered.
func (s *BaseDmSqlParserListener) EnterPlsql(ctx *PlsqlContext) {}

// ExitPlsql is called when production plsql is exited.
func (s *BaseDmSqlParserListener) ExitPlsql(ctx *PlsqlContext) {}

// EnterUr_option is called when production ur_option is entered.
func (s *BaseDmSqlParserListener) EnterUr_option(ctx *Ur_optionContext) {}

// ExitUr_option is called when production ur_option is exited.
func (s *BaseDmSqlParserListener) ExitUr_option(ctx *Ur_optionContext) {}

// EnterFlashback_trig_enable is called when production flashback_trig_enable is entered.
func (s *BaseDmSqlParserListener) EnterFlashback_trig_enable(ctx *Flashback_trig_enableContext) {}

// ExitFlashback_trig_enable is called when production flashback_trig_enable is exited.
func (s *BaseDmSqlParserListener) ExitFlashback_trig_enable(ctx *Flashback_trig_enableContext) {}

// EnterScn_or_lsn is called when production scn_or_lsn is entered.
func (s *BaseDmSqlParserListener) EnterScn_or_lsn(ctx *Scn_or_lsnContext) {}

// ExitScn_or_lsn is called when production scn_or_lsn is exited.
func (s *BaseDmSqlParserListener) ExitScn_or_lsn(ctx *Scn_or_lsnContext) {}

// EnterFull_table_name_list is called when production full_table_name_list is entered.
func (s *BaseDmSqlParserListener) EnterFull_table_name_list(ctx *Full_table_name_listContext) {}

// ExitFull_table_name_list is called when production full_table_name_list is exited.
func (s *BaseDmSqlParserListener) ExitFull_table_name_list(ctx *Full_table_name_listContext) {}

// EnterFlashback_tab_stmt is called when production flashback_tab_stmt is entered.
func (s *BaseDmSqlParserListener) EnterFlashback_tab_stmt(ctx *Flashback_tab_stmtContext) {}

// ExitFlashback_tab_stmt is called when production flashback_tab_stmt is exited.
func (s *BaseDmSqlParserListener) ExitFlashback_tab_stmt(ctx *Flashback_tab_stmtContext) {}

// EnterRename is called when production rename is entered.
func (s *BaseDmSqlParserListener) EnterRename(ctx *RenameContext) {}

// ExitRename is called when production rename is exited.
func (s *BaseDmSqlParserListener) ExitRename(ctx *RenameContext) {}

// EnterAlter_system_set_stmt is called when production alter_system_set_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_system_set_stmt(ctx *Alter_system_set_stmtContext) {}

// ExitAlter_system_set_stmt is called when production alter_system_set_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_system_set_stmt(ctx *Alter_system_set_stmtContext) {}

// EnterDefer is called when production defer is entered.
func (s *BaseDmSqlParserListener) EnterDefer(ctx *DeferContext) {}

// ExitDefer is called when production defer is exited.
func (s *BaseDmSqlParserListener) ExitDefer(ctx *DeferContext) {}

// EnterScope is called when production scope is entered.
func (s *BaseDmSqlParserListener) EnterScope(ctx *ScopeContext) {}

// ExitScope is called when production scope is exited.
func (s *BaseDmSqlParserListener) ExitScope(ctx *ScopeContext) {}

// EnterAlter_session_stmt is called when production alter_session_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_session_stmt(ctx *Alter_session_stmtContext) {}

// ExitAlter_session_stmt is called when production alter_session_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_session_stmt(ctx *Alter_session_stmtContext) {}

// EnterParallel_mode is called when production parallel_mode is entered.
func (s *BaseDmSqlParserListener) EnterParallel_mode(ctx *Parallel_modeContext) {}

// ExitParallel_mode is called when production parallel_mode is exited.
func (s *BaseDmSqlParserListener) ExitParallel_mode(ctx *Parallel_modeContext) {}

// EnterParallel_degree is called when production parallel_degree is entered.
func (s *BaseDmSqlParserListener) EnterParallel_degree(ctx *Parallel_degreeContext) {}

// ExitParallel_degree is called when production parallel_degree is exited.
func (s *BaseDmSqlParserListener) ExitParallel_degree(ctx *Parallel_degreeContext) {}

// EnterPurge is called when production purge is entered.
func (s *BaseDmSqlParserListener) EnterPurge(ctx *PurgeContext) {}

// ExitPurge is called when production purge is exited.
func (s *BaseDmSqlParserListener) ExitPurge(ctx *PurgeContext) {}

// EnterSess_id is called when production sess_id is entered.
func (s *BaseDmSqlParserListener) EnterSess_id(ctx *Sess_idContext) {}

// ExitSess_id is called when production sess_id is exited.
func (s *BaseDmSqlParserListener) ExitSess_id(ctx *Sess_idContext) {}

// EnterSet_time_zone_string is called when production set_time_zone_string is entered.
func (s *BaseDmSqlParserListener) EnterSet_time_zone_string(ctx *Set_time_zone_stringContext) {}

// ExitSet_time_zone_string is called when production set_time_zone_string is exited.
func (s *BaseDmSqlParserListener) ExitSet_time_zone_string(ctx *Set_time_zone_stringContext) {}

// EnterSess_attr is called when production sess_attr is entered.
func (s *BaseDmSqlParserListener) EnterSess_attr(ctx *Sess_attrContext) {}

// ExitSess_attr is called when production sess_attr is exited.
func (s *BaseDmSqlParserListener) ExitSess_attr(ctx *Sess_attrContext) {}

// EnterSess_attr_val is called when production sess_attr_val is entered.
func (s *BaseDmSqlParserListener) EnterSess_attr_val(ctx *Sess_attr_valContext) {}

// ExitSess_attr_val is called when production sess_attr_val is exited.
func (s *BaseDmSqlParserListener) ExitSess_attr_val(ctx *Sess_attr_valContext) {}

// EnterSet_schema_stmt is called when production set_schema_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSet_schema_stmt(ctx *Set_schema_stmtContext) {}

// ExitSet_schema_stmt is called when production set_schema_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSet_schema_stmt(ctx *Set_schema_stmtContext) {}

// EnterPlblock is called when production plblock is entered.
func (s *BaseDmSqlParserListener) EnterPlblock(ctx *PlblockContext) {}

// ExitPlblock is called when production plblock is exited.
func (s *BaseDmSqlParserListener) ExitPlblock(ctx *PlblockContext) {}

// EnterExcept_option is called when production except_option is entered.
func (s *BaseDmSqlParserListener) EnterExcept_option(ctx *Except_optionContext) {}

// ExitExcept_option is called when production except_option is exited.
func (s *BaseDmSqlParserListener) ExitExcept_option(ctx *Except_optionContext) {}

// EnterFinally_option is called when production finally_option is entered.
func (s *BaseDmSqlParserListener) EnterFinally_option(ctx *Finally_optionContext) {}

// ExitFinally_option is called when production finally_option is exited.
func (s *BaseDmSqlParserListener) ExitFinally_option(ctx *Finally_optionContext) {}

// EnterFinally_tail is called when production finally_tail is entered.
func (s *BaseDmSqlParserListener) EnterFinally_tail(ctx *Finally_tailContext) {}

// ExitFinally_tail is called when production finally_tail is exited.
func (s *BaseDmSqlParserListener) ExitFinally_tail(ctx *Finally_tailContext) {}

// EnterExcept_handler_list is called when production except_handler_list is entered.
func (s *BaseDmSqlParserListener) EnterExcept_handler_list(ctx *Except_handler_listContext) {}

// ExitExcept_handler_list is called when production except_handler_list is exited.
func (s *BaseDmSqlParserListener) ExitExcept_handler_list(ctx *Except_handler_listContext) {}

// EnterExcept_handler is called when production except_handler is entered.
func (s *BaseDmSqlParserListener) EnterExcept_handler(ctx *Except_handlerContext) {}

// ExitExcept_handler is called when production except_handler is exited.
func (s *BaseDmSqlParserListener) ExitExcept_handler(ctx *Except_handlerContext) {}

// EnterExcept_name is called when production except_name is entered.
func (s *BaseDmSqlParserListener) EnterExcept_name(ctx *Except_nameContext) {}

// ExitExcept_name is called when production except_name is exited.
func (s *BaseDmSqlParserListener) ExitExcept_name(ctx *Except_nameContext) {}

// EnterExcept_list is called when production except_list is entered.
func (s *BaseDmSqlParserListener) EnterExcept_list(ctx *Except_listContext) {}

// ExitExcept_list is called when production except_list is exited.
func (s *BaseDmSqlParserListener) ExitExcept_list(ctx *Except_listContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseDmSqlParserListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseDmSqlParserListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterIf_stmt_clause is called when production if_stmt_clause is entered.
func (s *BaseDmSqlParserListener) EnterIf_stmt_clause(ctx *If_stmt_clauseContext) {}

// ExitIf_stmt_clause is called when production if_stmt_clause is exited.
func (s *BaseDmSqlParserListener) ExitIf_stmt_clause(ctx *If_stmt_clauseContext) {}

// EnterIf_condition_clause is called when production if_condition_clause is entered.
func (s *BaseDmSqlParserListener) EnterIf_condition_clause(ctx *If_condition_clauseContext) {}

// ExitIf_condition_clause is called when production if_condition_clause is exited.
func (s *BaseDmSqlParserListener) ExitIf_condition_clause(ctx *If_condition_clauseContext) {}

// EnterIf_then_clause is called when production if_then_clause is entered.
func (s *BaseDmSqlParserListener) EnterIf_then_clause(ctx *If_then_clauseContext) {}

// ExitIf_then_clause is called when production if_then_clause is exited.
func (s *BaseDmSqlParserListener) ExitIf_then_clause(ctx *If_then_clauseContext) {}

// EnterElseif_lst_option is called when production elseif_lst_option is entered.
func (s *BaseDmSqlParserListener) EnterElseif_lst_option(ctx *Elseif_lst_optionContext) {}

// ExitElseif_lst_option is called when production elseif_lst_option is exited.
func (s *BaseDmSqlParserListener) ExitElseif_lst_option(ctx *Elseif_lst_optionContext) {}

// EnterElseif_clause is called when production elseif_clause is entered.
func (s *BaseDmSqlParserListener) EnterElseif_clause(ctx *Elseif_clauseContext) {}

// ExitElseif_clause is called when production elseif_clause is exited.
func (s *BaseDmSqlParserListener) ExitElseif_clause(ctx *Elseif_clauseContext) {}

// EnterElse_option is called when production else_option is entered.
func (s *BaseDmSqlParserListener) EnterElse_option(ctx *Else_optionContext) {}

// ExitElse_option is called when production else_option is exited.
func (s *BaseDmSqlParserListener) ExitElse_option(ctx *Else_optionContext) {}

// EnterSs_if_stmt_clause is called when production ss_if_stmt_clause is entered.
func (s *BaseDmSqlParserListener) EnterSs_if_stmt_clause(ctx *Ss_if_stmt_clauseContext) {}

// ExitSs_if_stmt_clause is called when production ss_if_stmt_clause is exited.
func (s *BaseDmSqlParserListener) ExitSs_if_stmt_clause(ctx *Ss_if_stmt_clauseContext) {}

// EnterSs_elseif_lst_option is called when production ss_elseif_lst_option is entered.
func (s *BaseDmSqlParserListener) EnterSs_elseif_lst_option(ctx *Ss_elseif_lst_optionContext) {}

// ExitSs_elseif_lst_option is called when production ss_elseif_lst_option is exited.
func (s *BaseDmSqlParserListener) ExitSs_elseif_lst_option(ctx *Ss_elseif_lst_optionContext) {}

// EnterSs_elseif_clause is called when production ss_elseif_clause is entered.
func (s *BaseDmSqlParserListener) EnterSs_elseif_clause(ctx *Ss_elseif_clauseContext) {}

// ExitSs_elseif_clause is called when production ss_elseif_clause is exited.
func (s *BaseDmSqlParserListener) ExitSs_elseif_clause(ctx *Ss_elseif_clauseContext) {}

// EnterSs_else_option is called when production ss_else_option is entered.
func (s *BaseDmSqlParserListener) EnterSs_else_option(ctx *Ss_else_optionContext) {}

// ExitSs_else_option is called when production ss_else_option is exited.
func (s *BaseDmSqlParserListener) ExitSs_else_option(ctx *Ss_else_optionContext) {}

// EnterCase_stmt is called when production case_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCase_stmt(ctx *Case_stmtContext) {}

// ExitCase_stmt is called when production case_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCase_stmt(ctx *Case_stmtContext) {}

// EnterPlsearched_when_clause is called when production plsearched_when_clause is entered.
func (s *BaseDmSqlParserListener) EnterPlsearched_when_clause(ctx *Plsearched_when_clauseContext) {}

// ExitPlsearched_when_clause is called when production plsearched_when_clause is exited.
func (s *BaseDmSqlParserListener) ExitPlsearched_when_clause(ctx *Plsearched_when_clauseContext) {}

// EnterPlsearched_when_list is called when production plsearched_when_list is entered.
func (s *BaseDmSqlParserListener) EnterPlsearched_when_list(ctx *Plsearched_when_listContext) {}

// ExitPlsearched_when_list is called when production plsearched_when_list is exited.
func (s *BaseDmSqlParserListener) ExitPlsearched_when_list(ctx *Plsearched_when_listContext) {}

// EnterCase_option is called when production case_option is entered.
func (s *BaseDmSqlParserListener) EnterCase_option(ctx *Case_optionContext) {}

// ExitCase_option is called when production case_option is exited.
func (s *BaseDmSqlParserListener) ExitCase_option(ctx *Case_optionContext) {}

// EnterAssign_stmt is called when production assign_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAssign_stmt(ctx *Assign_stmtContext) {}

// ExitAssign_stmt is called when production assign_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAssign_stmt(ctx *Assign_stmtContext) {}

// EnterAssign_obj is called when production assign_obj is entered.
func (s *BaseDmSqlParserListener) EnterAssign_obj(ctx *Assign_objContext) {}

// ExitAssign_obj is called when production assign_obj is exited.
func (s *BaseDmSqlParserListener) ExitAssign_obj(ctx *Assign_objContext) {}

// EnterAssign_obj2 is called when production assign_obj2 is entered.
func (s *BaseDmSqlParserListener) EnterAssign_obj2(ctx *Assign_obj2Context) {}

// ExitAssign_obj2 is called when production assign_obj2 is exited.
func (s *BaseDmSqlParserListener) ExitAssign_obj2(ctx *Assign_obj2Context) {}

// EnterAssign_op is called when production assign_op is entered.
func (s *BaseDmSqlParserListener) EnterAssign_op(ctx *Assign_opContext) {}

// ExitAssign_op is called when production assign_op is exited.
func (s *BaseDmSqlParserListener) ExitAssign_op(ctx *Assign_opContext) {}

// EnterGoto_stmt is called when production goto_stmt is entered.
func (s *BaseDmSqlParserListener) EnterGoto_stmt(ctx *Goto_stmtContext) {}

// ExitGoto_stmt is called when production goto_stmt is exited.
func (s *BaseDmSqlParserListener) ExitGoto_stmt(ctx *Goto_stmtContext) {}

// EnterWhile_stmt is called when production while_stmt is entered.
func (s *BaseDmSqlParserListener) EnterWhile_stmt(ctx *While_stmtContext) {}

// ExitWhile_stmt is called when production while_stmt is exited.
func (s *BaseDmSqlParserListener) ExitWhile_stmt(ctx *While_stmtContext) {}

// EnterLoop_stmt is called when production loop_stmt is entered.
func (s *BaseDmSqlParserListener) EnterLoop_stmt(ctx *Loop_stmtContext) {}

// ExitLoop_stmt is called when production loop_stmt is exited.
func (s *BaseDmSqlParserListener) ExitLoop_stmt(ctx *Loop_stmtContext) {}

// EnterRepeat_stmt is called when production repeat_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRepeat_stmt(ctx *Repeat_stmtContext) {}

// ExitRepeat_stmt is called when production repeat_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRepeat_stmt(ctx *Repeat_stmtContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseDmSqlParserListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseDmSqlParserListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterForall_stmt is called when production forall_stmt is entered.
func (s *BaseDmSqlParserListener) EnterForall_stmt(ctx *Forall_stmtContext) {}

// ExitForall_stmt is called when production forall_stmt is exited.
func (s *BaseDmSqlParserListener) ExitForall_stmt(ctx *Forall_stmtContext) {}

// EnterForall_between_option is called when production forall_between_option is entered.
func (s *BaseDmSqlParserListener) EnterForall_between_option(ctx *Forall_between_optionContext) {}

// ExitForall_between_option is called when production forall_between_option is exited.
func (s *BaseDmSqlParserListener) ExitForall_between_option(ctx *Forall_between_optionContext) {}

// EnterForall_save_exception_option is called when production forall_save_exception_option is entered.
func (s *BaseDmSqlParserListener) EnterForall_save_exception_option(ctx *Forall_save_exception_optionContext) {
}

// ExitForall_save_exception_option is called when production forall_save_exception_option is exited.
func (s *BaseDmSqlParserListener) ExitForall_save_exception_option(ctx *Forall_save_exception_optionContext) {
}

// EnterForall_index_values is called when production forall_index_values is entered.
func (s *BaseDmSqlParserListener) EnterForall_index_values(ctx *Forall_index_valuesContext) {}

// ExitForall_index_values is called when production forall_index_values is exited.
func (s *BaseDmSqlParserListener) ExitForall_index_values(ctx *Forall_index_valuesContext) {}

// EnterForall_start is called when production forall_start is entered.
func (s *BaseDmSqlParserListener) EnterForall_start(ctx *Forall_startContext) {}

// ExitForall_start is called when production forall_start is exited.
func (s *BaseDmSqlParserListener) ExitForall_start(ctx *Forall_startContext) {}

// EnterForall_dml_stmt is called when production forall_dml_stmt is entered.
func (s *BaseDmSqlParserListener) EnterForall_dml_stmt(ctx *Forall_dml_stmtContext) {}

// ExitForall_dml_stmt is called when production forall_dml_stmt is exited.
func (s *BaseDmSqlParserListener) ExitForall_dml_stmt(ctx *Forall_dml_stmtContext) {}

// EnterIn_option is called when production in_option is entered.
func (s *BaseDmSqlParserListener) EnterIn_option(ctx *In_optionContext) {}

// ExitIn_option is called when production in_option is exited.
func (s *BaseDmSqlParserListener) ExitIn_option(ctx *In_optionContext) {}

// EnterFor_condition is called when production for_condition is entered.
func (s *BaseDmSqlParserListener) EnterFor_condition(ctx *For_conditionContext) {}

// ExitFor_condition is called when production for_condition is exited.
func (s *BaseDmSqlParserListener) ExitFor_condition(ctx *For_conditionContext) {}

// EnterPipe_row_stmt is called when production pipe_row_stmt is entered.
func (s *BaseDmSqlParserListener) EnterPipe_row_stmt(ctx *Pipe_row_stmtContext) {}

// ExitPipe_row_stmt is called when production pipe_row_stmt is exited.
func (s *BaseDmSqlParserListener) ExitPipe_row_stmt(ctx *Pipe_row_stmtContext) {}

// EnterExit_stmt is called when production exit_stmt is entered.
func (s *BaseDmSqlParserListener) EnterExit_stmt(ctx *Exit_stmtContext) {}

// ExitExit_stmt is called when production exit_stmt is exited.
func (s *BaseDmSqlParserListener) ExitExit_stmt(ctx *Exit_stmtContext) {}

// EnterContinue_stmt is called when production continue_stmt is entered.
func (s *BaseDmSqlParserListener) EnterContinue_stmt(ctx *Continue_stmtContext) {}

// ExitContinue_stmt is called when production continue_stmt is exited.
func (s *BaseDmSqlParserListener) ExitContinue_stmt(ctx *Continue_stmtContext) {}

// EnterNull_stmt is called when production null_stmt is entered.
func (s *BaseDmSqlParserListener) EnterNull_stmt(ctx *Null_stmtContext) {}

// ExitNull_stmt is called when production null_stmt is exited.
func (s *BaseDmSqlParserListener) ExitNull_stmt(ctx *Null_stmtContext) {}

// EnterPrint_stmt is called when production print_stmt is entered.
func (s *BaseDmSqlParserListener) EnterPrint_stmt(ctx *Print_stmtContext) {}

// ExitPrint_stmt is called when production print_stmt is exited.
func (s *BaseDmSqlParserListener) ExitPrint_stmt(ctx *Print_stmtContext) {}

// EnterExecute_stmt is called when production execute_stmt is entered.
func (s *BaseDmSqlParserListener) EnterExecute_stmt(ctx *Execute_stmtContext) {}

// ExitExecute_stmt is called when production execute_stmt is exited.
func (s *BaseDmSqlParserListener) ExitExecute_stmt(ctx *Execute_stmtContext) {}

// EnterDyn_return is called when production dyn_return is entered.
func (s *BaseDmSqlParserListener) EnterDyn_return(ctx *Dyn_returnContext) {}

// ExitDyn_return is called when production dyn_return is exited.
func (s *BaseDmSqlParserListener) ExitDyn_return(ctx *Dyn_returnContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseDmSqlParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseDmSqlParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterUsing_exp_list is called when production using_exp_list is entered.
func (s *BaseDmSqlParserListener) EnterUsing_exp_list(ctx *Using_exp_listContext) {}

// ExitUsing_exp_list is called when production using_exp_list is exited.
func (s *BaseDmSqlParserListener) ExitUsing_exp_list(ctx *Using_exp_listContext) {}

// EnterUsing_exp is called when production using_exp is entered.
func (s *BaseDmSqlParserListener) EnterUsing_exp(ctx *Using_expContext) {}

// ExitUsing_exp is called when production using_exp is exited.
func (s *BaseDmSqlParserListener) ExitUsing_exp(ctx *Using_expContext) {}

// EnterAlter_proc_stmt is called when production alter_proc_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_proc_stmt(ctx *Alter_proc_stmtContext) {}

// ExitAlter_proc_stmt is called when production alter_proc_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_proc_stmt(ctx *Alter_proc_stmtContext) {}

// EnterAlter_func_stmt is called when production alter_func_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_func_stmt(ctx *Alter_func_stmtContext) {}

// ExitAlter_func_stmt is called when production alter_func_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_func_stmt(ctx *Alter_func_stmtContext) {}

// EnterAlter_package_stmt is called when production alter_package_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_package_stmt(ctx *Alter_package_stmtContext) {}

// ExitAlter_package_stmt is called when production alter_package_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_package_stmt(ctx *Alter_package_stmtContext) {}

// EnterPkg_type is called when production pkg_type is entered.
func (s *BaseDmSqlParserListener) EnterPkg_type(ctx *Pkg_typeContext) {}

// ExitPkg_type is called when production pkg_type is exited.
func (s *BaseDmSqlParserListener) ExitPkg_type(ctx *Pkg_typeContext) {}

// EnterDeclare_opt is called when production declare_opt is entered.
func (s *BaseDmSqlParserListener) EnterDeclare_opt(ctx *Declare_optContext) {}

// ExitDeclare_opt is called when production declare_opt is exited.
func (s *BaseDmSqlParserListener) ExitDeclare_opt(ctx *Declare_optContext) {}

// EnterAlter_table_stmt is called when production alter_table_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// ExitAlter_table_stmt is called when production alter_table_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_table_stmt(ctx *Alter_table_stmtContext) {}

// EnterAlter_tag is called when production alter_tag is entered.
func (s *BaseDmSqlParserListener) EnterAlter_tag(ctx *Alter_tagContext) {}

// ExitAlter_tag is called when production alter_tag is exited.
func (s *BaseDmSqlParserListener) ExitAlter_tag(ctx *Alter_tagContext) {}

// EnterAlter_index_stmt is called when production alter_index_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_index_stmt(ctx *Alter_index_stmtContext) {}

// ExitAlter_index_stmt is called when production alter_index_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_index_stmt(ctx *Alter_index_stmtContext) {}

// EnterFull_index_name is called when production full_index_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_index_name(ctx *Full_index_nameContext) {}

// ExitFull_index_name is called when production full_index_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_index_name(ctx *Full_index_nameContext) {}

// EnterAlter_index_action is called when production alter_index_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_index_action(ctx *Alter_index_actionContext) {}

// ExitAlter_index_action is called when production alter_index_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_index_action(ctx *Alter_index_actionContext) {}

// EnterRebuild_clause is called when production rebuild_clause is entered.
func (s *BaseDmSqlParserListener) EnterRebuild_clause(ctx *Rebuild_clauseContext) {}

// ExitRebuild_clause is called when production rebuild_clause is exited.
func (s *BaseDmSqlParserListener) ExitRebuild_clause(ctx *Rebuild_clauseContext) {}

// EnterExclusive_options is called when production exclusive_options is entered.
func (s *BaseDmSqlParserListener) EnterExclusive_options(ctx *Exclusive_optionsContext) {}

// ExitExclusive_options is called when production exclusive_options is exited.
func (s *BaseDmSqlParserListener) ExitExclusive_options(ctx *Exclusive_optionsContext) {}

// EnterAsynchronous_options is called when production asynchronous_options is entered.
func (s *BaseDmSqlParserListener) EnterAsynchronous_options(ctx *Asynchronous_optionsContext) {}

// ExitAsynchronous_options is called when production asynchronous_options is exited.
func (s *BaseDmSqlParserListener) ExitAsynchronous_options(ctx *Asynchronous_optionsContext) {}

// EnterVisible_clause is called when production visible_clause is entered.
func (s *BaseDmSqlParserListener) EnterVisible_clause(ctx *Visible_clauseContext) {}

// ExitVisible_clause is called when production visible_clause is exited.
func (s *BaseDmSqlParserListener) ExitVisible_clause(ctx *Visible_clauseContext) {}

// EnterColumn_def_list is called when production column_def_list is entered.
func (s *BaseDmSqlParserListener) EnterColumn_def_list(ctx *Column_def_listContext) {}

// ExitColumn_def_list is called when production column_def_list is exited.
func (s *BaseDmSqlParserListener) ExitColumn_def_list(ctx *Column_def_listContext) {}

// EnterLock is called when production lock is entered.
func (s *BaseDmSqlParserListener) EnterLock(ctx *LockContext) {}

// ExitLock is called when production lock is exited.
func (s *BaseDmSqlParserListener) ExitLock(ctx *LockContext) {}

// EnterAlter_table_partition_action is called when production alter_table_partition_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_table_partition_action(ctx *Alter_table_partition_actionContext) {
}

// ExitAlter_table_partition_action is called when production alter_table_partition_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_table_partition_action(ctx *Alter_table_partition_actionContext) {
}

// EnterTemplate_info is called when production template_info is entered.
func (s *BaseDmSqlParserListener) EnterTemplate_info(ctx *Template_infoContext) {}

// ExitTemplate_info is called when production template_info is exited.
func (s *BaseDmSqlParserListener) ExitTemplate_info(ctx *Template_infoContext) {}

// EnterTemplate_item_2 is called when production template_item_2 is entered.
func (s *BaseDmSqlParserListener) EnterTemplate_item_2(ctx *Template_item_2Context) {}

// ExitTemplate_item_2 is called when production template_item_2 is exited.
func (s *BaseDmSqlParserListener) ExitTemplate_item_2(ctx *Template_item_2Context) {}

// EnterTemplate_item_1 is called when production template_item_1 is entered.
func (s *BaseDmSqlParserListener) EnterTemplate_item_1(ctx *Template_item_1Context) {}

// ExitTemplate_item_1 is called when production template_item_1 is exited.
func (s *BaseDmSqlParserListener) ExitTemplate_item_1(ctx *Template_item_1Context) {}

// EnterIncluding_indexes is called when production including_indexes is entered.
func (s *BaseDmSqlParserListener) EnterIncluding_indexes(ctx *Including_indexesContext) {}

// ExitIncluding_indexes is called when production including_indexes is exited.
func (s *BaseDmSqlParserListener) ExitIncluding_indexes(ctx *Including_indexesContext) {}

// EnterTruncate_partition_name is called when production truncate_partition_name is entered.
func (s *BaseDmSqlParserListener) EnterTruncate_partition_name(ctx *Truncate_partition_nameContext) {}

// ExitTruncate_partition_name is called when production truncate_partition_name is exited.
func (s *BaseDmSqlParserListener) ExitTruncate_partition_name(ctx *Truncate_partition_nameContext) {}

// EnterCons_enable is called when production cons_enable is entered.
func (s *BaseDmSqlParserListener) EnterCons_enable(ctx *Cons_enableContext) {}

// ExitCons_enable is called when production cons_enable is exited.
func (s *BaseDmSqlParserListener) ExitCons_enable(ctx *Cons_enableContext) {}

// EnterReuse_storage_option is called when production reuse_storage_option is entered.
func (s *BaseDmSqlParserListener) EnterReuse_storage_option(ctx *Reuse_storage_optionContext) {}

// ExitReuse_storage_option is called when production reuse_storage_option is exited.
func (s *BaseDmSqlParserListener) ExitReuse_storage_option(ctx *Reuse_storage_optionContext) {}

// EnterAlter_table_action is called when production alter_table_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_table_action(ctx *Alter_table_actionContext) {}

// ExitAlter_table_action is called when production alter_table_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_table_action(ctx *Alter_table_actionContext) {}

// EnterFast_flag is called when production fast_flag is entered.
func (s *BaseDmSqlParserListener) EnterFast_flag(ctx *Fast_flagContext) {}

// ExitFast_flag is called when production fast_flag is exited.
func (s *BaseDmSqlParserListener) ExitFast_flag(ctx *Fast_flagContext) {}

// EnterStorage_stat_flag is called when production storage_stat_flag is entered.
func (s *BaseDmSqlParserListener) EnterStorage_stat_flag(ctx *Storage_stat_flagContext) {}

// ExitStorage_stat_flag is called when production storage_stat_flag is exited.
func (s *BaseDmSqlParserListener) ExitStorage_stat_flag(ctx *Storage_stat_flagContext) {}

// EnterStorage_stat_cols is called when production storage_stat_cols is entered.
func (s *BaseDmSqlParserListener) EnterStorage_stat_cols(ctx *Storage_stat_colsContext) {}

// ExitStorage_stat_cols is called when production storage_stat_cols is exited.
func (s *BaseDmSqlParserListener) ExitStorage_stat_cols(ctx *Storage_stat_colsContext) {}

// EnterHfs_rebuild_level is called when production hfs_rebuild_level is entered.
func (s *BaseDmSqlParserListener) EnterHfs_rebuild_level(ctx *Hfs_rebuild_levelContext) {}

// ExitHfs_rebuild_level is called when production hfs_rebuild_level is exited.
func (s *BaseDmSqlParserListener) ExitHfs_rebuild_level(ctx *Hfs_rebuild_levelContext) {}

// EnterAta_lock_option is called when production ata_lock_option is entered.
func (s *BaseDmSqlParserListener) EnterAta_lock_option(ctx *Ata_lock_optionContext) {}

// ExitAta_lock_option is called when production ata_lock_option is exited.
func (s *BaseDmSqlParserListener) ExitAta_lock_option(ctx *Ata_lock_optionContext) {}

// EnterMdf_column_def_list is called when production mdf_column_def_list is entered.
func (s *BaseDmSqlParserListener) EnterMdf_column_def_list(ctx *Mdf_column_def_listContext) {}

// ExitMdf_column_def_list is called when production mdf_column_def_list is exited.
func (s *BaseDmSqlParserListener) ExitMdf_column_def_list(ctx *Mdf_column_def_listContext) {}

// EnterMdf_column_def is called when production mdf_column_def is entered.
func (s *BaseDmSqlParserListener) EnterMdf_column_def(ctx *Mdf_column_defContext) {}

// ExitMdf_column_def is called when production mdf_column_def is exited.
func (s *BaseDmSqlParserListener) ExitMdf_column_def(ctx *Mdf_column_defContext) {}

// EnterColumn_def is called when production column_def is entered.
func (s *BaseDmSqlParserListener) EnterColumn_def(ctx *Column_defContext) {}

// ExitColumn_def is called when production column_def is exited.
func (s *BaseDmSqlParserListener) ExitColumn_def(ctx *Column_defContext) {}

// EnterColumn_def_ex is called when production column_def_ex is entered.
func (s *BaseDmSqlParserListener) EnterColumn_def_ex(ctx *Column_def_exContext) {}

// ExitColumn_def_ex is called when production column_def_ex is exited.
func (s *BaseDmSqlParserListener) ExitColumn_def_ex(ctx *Column_def_exContext) {}

// EnterColumn_def_low is called when production column_def_low is entered.
func (s *BaseDmSqlParserListener) EnterColumn_def_low(ctx *Column_def_lowContext) {}

// ExitColumn_def_low is called when production column_def_low is exited.
func (s *BaseDmSqlParserListener) ExitColumn_def_low(ctx *Column_def_lowContext) {}

// EnterVirtual_column_datatype is called when production virtual_column_datatype is entered.
func (s *BaseDmSqlParserListener) EnterVirtual_column_datatype(ctx *Virtual_column_datatypeContext) {}

// ExitVirtual_column_datatype is called when production virtual_column_datatype is exited.
func (s *BaseDmSqlParserListener) ExitVirtual_column_datatype(ctx *Virtual_column_datatypeContext) {}

// EnterVirtual_column_generated is called when production virtual_column_generated is entered.
func (s *BaseDmSqlParserListener) EnterVirtual_column_generated(ctx *Virtual_column_generatedContext) {
}

// ExitVirtual_column_generated is called when production virtual_column_generated is exited.
func (s *BaseDmSqlParserListener) ExitVirtual_column_generated(ctx *Virtual_column_generatedContext) {
}

// EnterVirtual_column_virtual is called when production virtual_column_virtual is entered.
func (s *BaseDmSqlParserListener) EnterVirtual_column_virtual(ctx *Virtual_column_virtualContext) {}

// ExitVirtual_column_virtual is called when production virtual_column_virtual is exited.
func (s *BaseDmSqlParserListener) ExitVirtual_column_virtual(ctx *Virtual_column_virtualContext) {}

// EnterVirtual_column_visible is called when production virtual_column_visible is entered.
func (s *BaseDmSqlParserListener) EnterVirtual_column_visible(ctx *Virtual_column_visibleContext) {}

// ExitVirtual_column_visible is called when production virtual_column_visible is exited.
func (s *BaseDmSqlParserListener) ExitVirtual_column_visible(ctx *Virtual_column_visibleContext) {}

// EnterVirtual_column_def is called when production virtual_column_def is entered.
func (s *BaseDmSqlParserListener) EnterVirtual_column_def(ctx *Virtual_column_defContext) {}

// ExitVirtual_column_def is called when production virtual_column_def is exited.
func (s *BaseDmSqlParserListener) ExitVirtual_column_def(ctx *Virtual_column_defContext) {}

// EnterCharset_option is called when production charset_option is entered.
func (s *BaseDmSqlParserListener) EnterCharset_option(ctx *Charset_optionContext) {}

// ExitCharset_option is called when production charset_option is exited.
func (s *BaseDmSqlParserListener) ExitCharset_option(ctx *Charset_optionContext) {}

// EnterColumn_def_4_option is called when production column_def_4_option is entered.
func (s *BaseDmSqlParserListener) EnterColumn_def_4_option(ctx *Column_def_4_optionContext) {}

// ExitColumn_def_4_option is called when production column_def_4_option is exited.
func (s *BaseDmSqlParserListener) ExitColumn_def_4_option(ctx *Column_def_4_optionContext) {}

// EnterAuto_update_clause is called when production auto_update_clause is entered.
func (s *BaseDmSqlParserListener) EnterAuto_update_clause(ctx *Auto_update_clauseContext) {}

// ExitAuto_update_clause is called when production auto_update_clause is exited.
func (s *BaseDmSqlParserListener) ExitAuto_update_clause(ctx *Auto_update_clauseContext) {}

// EnterUpdate_exp is called when production update_exp is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_exp(ctx *Update_expContext) {}

// ExitUpdate_exp is called when production update_exp is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_exp(ctx *Update_expContext) {}

// EnterIdentity_clause is called when production identity_clause is entered.
func (s *BaseDmSqlParserListener) EnterIdentity_clause(ctx *Identity_clauseContext) {}

// ExitIdentity_clause is called when production identity_clause is exited.
func (s *BaseDmSqlParserListener) ExitIdentity_clause(ctx *Identity_clauseContext) {}

// EnterDefault_clause_with_on_null_opt is called when production default_clause_with_on_null_opt is entered.
func (s *BaseDmSqlParserListener) EnterDefault_clause_with_on_null_opt(ctx *Default_clause_with_on_null_optContext) {
}

// ExitDefault_clause_with_on_null_opt is called when production default_clause_with_on_null_opt is exited.
func (s *BaseDmSqlParserListener) ExitDefault_clause_with_on_null_opt(ctx *Default_clause_with_on_null_optContext) {
}

// EnterDefault_clause is called when production default_clause is entered.
func (s *BaseDmSqlParserListener) EnterDefault_clause(ctx *Default_clauseContext) {}

// ExitDefault_clause is called when production default_clause is exited.
func (s *BaseDmSqlParserListener) ExitDefault_clause(ctx *Default_clauseContext) {}

// EnterDefault_exp is called when production default_exp is entered.
func (s *BaseDmSqlParserListener) EnterDefault_exp(ctx *Default_expContext) {}

// ExitDefault_exp is called when production default_exp is exited.
func (s *BaseDmSqlParserListener) ExitDefault_exp(ctx *Default_expContext) {}

// EnterColumn_constraint_def is called when production column_constraint_def is entered.
func (s *BaseDmSqlParserListener) EnterColumn_constraint_def(ctx *Column_constraint_defContext) {}

// ExitColumn_constraint_def is called when production column_constraint_def is exited.
func (s *BaseDmSqlParserListener) ExitColumn_constraint_def(ctx *Column_constraint_defContext) {}

// EnterConstraint_name_def_options is called when production constraint_name_def_options is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_name_def_options(ctx *Constraint_name_def_optionsContext) {
}

// ExitConstraint_name_def_options is called when production constraint_name_def_options is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_name_def_options(ctx *Constraint_name_def_optionsContext) {
}

// EnterConstraint_name_def is called when production constraint_name_def is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_name_def(ctx *Constraint_name_defContext) {}

// ExitConstraint_name_def is called when production constraint_name_def is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_name_def(ctx *Constraint_name_defContext) {}

// EnterColumn_constraints is called when production column_constraints is entered.
func (s *BaseDmSqlParserListener) EnterColumn_constraints(ctx *Column_constraintsContext) {}

// ExitColumn_constraints is called when production column_constraints is exited.
func (s *BaseDmSqlParserListener) ExitColumn_constraints(ctx *Column_constraintsContext) {}

// EnterColumn_constraint is called when production column_constraint is entered.
func (s *BaseDmSqlParserListener) EnterColumn_constraint(ctx *Column_constraintContext) {}

// ExitColumn_constraint is called when production column_constraint is exited.
func (s *BaseDmSqlParserListener) ExitColumn_constraint(ctx *Column_constraintContext) {}

// EnterColumn_constraint_action is called when production column_constraint_action is entered.
func (s *BaseDmSqlParserListener) EnterColumn_constraint_action(ctx *Column_constraint_actionContext) {
}

// ExitColumn_constraint_action is called when production column_constraint_action is exited.
func (s *BaseDmSqlParserListener) ExitColumn_constraint_action(ctx *Column_constraint_actionContext) {
}

// EnterNot_null_spec is called when production not_null_spec is entered.
func (s *BaseDmSqlParserListener) EnterNot_null_spec(ctx *Not_null_specContext) {}

// ExitNot_null_spec is called when production not_null_spec is exited.
func (s *BaseDmSqlParserListener) ExitNot_null_spec(ctx *Not_null_specContext) {}

// EnterUnique_spec is called when production unique_spec is entered.
func (s *BaseDmSqlParserListener) EnterUnique_spec(ctx *Unique_specContext) {}

// ExitUnique_spec is called when production unique_spec is exited.
func (s *BaseDmSqlParserListener) ExitUnique_spec(ctx *Unique_specContext) {}

// EnterRefs_spec is called when production refs_spec is entered.
func (s *BaseDmSqlParserListener) EnterRefs_spec(ctx *Refs_specContext) {}

// ExitRefs_spec is called when production refs_spec is exited.
func (s *BaseDmSqlParserListener) ExitRefs_spec(ctx *Refs_specContext) {}

// EnterRefs_spec_action is called when production refs_spec_action is entered.
func (s *BaseDmSqlParserListener) EnterRefs_spec_action(ctx *Refs_spec_actionContext) {}

// ExitRefs_spec_action is called when production refs_spec_action is exited.
func (s *BaseDmSqlParserListener) ExitRefs_spec_action(ctx *Refs_spec_actionContext) {}

// EnterForeign_key is called when production foreign_key is entered.
func (s *BaseDmSqlParserListener) EnterForeign_key(ctx *Foreign_keyContext) {}

// ExitForeign_key is called when production foreign_key is exited.
func (s *BaseDmSqlParserListener) ExitForeign_key(ctx *Foreign_keyContext) {}

// EnterRefd_table_and_columns is called when production refd_table_and_columns is entered.
func (s *BaseDmSqlParserListener) EnterRefd_table_and_columns(ctx *Refd_table_and_columnsContext) {}

// ExitRefd_table_and_columns is called when production refd_table_and_columns is exited.
func (s *BaseDmSqlParserListener) ExitRefd_table_and_columns(ctx *Refd_table_and_columnsContext) {}

// EnterRef_column_list is called when production ref_column_list is entered.
func (s *BaseDmSqlParserListener) EnterRef_column_list(ctx *Ref_column_listContext) {}

// ExitRef_column_list is called when production ref_column_list is exited.
func (s *BaseDmSqlParserListener) ExitRef_column_list(ctx *Ref_column_listContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterColumn_list2 is called when production column_list2 is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list2(ctx *Column_list2Context) {}

// ExitColumn_list2 is called when production column_list2 is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list2(ctx *Column_list2Context) {}

// EnterFull_column_list is called when production full_column_list is entered.
func (s *BaseDmSqlParserListener) EnterFull_column_list(ctx *Full_column_listContext) {}

// ExitFull_column_list is called when production full_column_list is exited.
func (s *BaseDmSqlParserListener) ExitFull_column_list(ctx *Full_column_listContext) {}

// EnterColumn_list_list is called when production column_list_list is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list_list(ctx *Column_list_listContext) {}

// ExitColumn_list_list is called when production column_list_list is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list_list(ctx *Column_list_listContext) {}

// EnterDrop_column_list is called when production drop_column_list is entered.
func (s *BaseDmSqlParserListener) EnterDrop_column_list(ctx *Drop_column_listContext) {}

// ExitDrop_column_list is called when production drop_column_list is exited.
func (s *BaseDmSqlParserListener) ExitDrop_column_list(ctx *Drop_column_listContext) {}

// EnterMatch_option is called when production match_option is entered.
func (s *BaseDmSqlParserListener) EnterMatch_option(ctx *Match_optionContext) {}

// ExitMatch_option is called when production match_option is exited.
func (s *BaseDmSqlParserListener) ExitMatch_option(ctx *Match_optionContext) {}

// EnterMatch_type is called when production match_type is entered.
func (s *BaseDmSqlParserListener) EnterMatch_type(ctx *Match_typeContext) {}

// ExitMatch_type is called when production match_type is exited.
func (s *BaseDmSqlParserListener) ExitMatch_type(ctx *Match_typeContext) {}

// EnterRef_triggered_action is called when production ref_triggered_action is entered.
func (s *BaseDmSqlParserListener) EnterRef_triggered_action(ctx *Ref_triggered_actionContext) {}

// ExitRef_triggered_action is called when production ref_triggered_action is exited.
func (s *BaseDmSqlParserListener) ExitRef_triggered_action(ctx *Ref_triggered_actionContext) {}

// EnterUpdate_rule is called when production update_rule is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_rule(ctx *Update_ruleContext) {}

// ExitUpdate_rule is called when production update_rule is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_rule(ctx *Update_ruleContext) {}

// EnterDelete_rule is called when production delete_rule is entered.
func (s *BaseDmSqlParserListener) EnterDelete_rule(ctx *Delete_ruleContext) {}

// ExitDelete_rule is called when production delete_rule is exited.
func (s *BaseDmSqlParserListener) ExitDelete_rule(ctx *Delete_ruleContext) {}

// EnterRef_action is called when production ref_action is entered.
func (s *BaseDmSqlParserListener) EnterRef_action(ctx *Ref_actionContext) {}

// ExitRef_action is called when production ref_action is exited.
func (s *BaseDmSqlParserListener) ExitRef_action(ctx *Ref_actionContext) {}

// EnterCheck_constraint_def is called when production check_constraint_def is entered.
func (s *BaseDmSqlParserListener) EnterCheck_constraint_def(ctx *Check_constraint_defContext) {}

// ExitCheck_constraint_def is called when production check_constraint_def is exited.
func (s *BaseDmSqlParserListener) ExitCheck_constraint_def(ctx *Check_constraint_defContext) {}

// EnterCheck_condition is called when production check_condition is entered.
func (s *BaseDmSqlParserListener) EnterCheck_condition(ctx *Check_conditionContext) {}

// ExitCheck_condition is called when production check_condition is exited.
func (s *BaseDmSqlParserListener) ExitCheck_condition(ctx *Check_conditionContext) {}

// EnterRestrict_cascade is called when production restrict_cascade is entered.
func (s *BaseDmSqlParserListener) EnterRestrict_cascade(ctx *Restrict_cascadeContext) {}

// ExitRestrict_cascade is called when production restrict_cascade is exited.
func (s *BaseDmSqlParserListener) ExitRestrict_cascade(ctx *Restrict_cascadeContext) {}

// EnterCascade_opt is called when production cascade_opt is entered.
func (s *BaseDmSqlParserListener) EnterCascade_opt(ctx *Cascade_optContext) {}

// ExitCascade_opt is called when production cascade_opt is exited.
func (s *BaseDmSqlParserListener) ExitCascade_opt(ctx *Cascade_optContext) {}

// EnterConstraint_name_options is called when production constraint_name_options is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_name_options(ctx *Constraint_name_optionsContext) {}

// ExitConstraint_name_options is called when production constraint_name_options is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_name_options(ctx *Constraint_name_optionsContext) {}

// EnterCheck_option_def_true is called when production check_option_def_true is entered.
func (s *BaseDmSqlParserListener) EnterCheck_option_def_true(ctx *Check_option_def_trueContext) {}

// ExitCheck_option_def_true is called when production check_option_def_true is exited.
func (s *BaseDmSqlParserListener) ExitCheck_option_def_true(ctx *Check_option_def_trueContext) {}

// EnterConstraint_attributes_options is called when production constraint_attributes_options is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_attributes_options(ctx *Constraint_attributes_optionsContext) {
}

// ExitConstraint_attributes_options is called when production constraint_attributes_options is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_attributes_options(ctx *Constraint_attributes_optionsContext) {
}

// EnterConstraint_attributes is called when production constraint_attributes is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_attributes(ctx *Constraint_attributesContext) {}

// ExitConstraint_attributes is called when production constraint_attributes is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_attributes(ctx *Constraint_attributesContext) {}

// EnterDeferrable_option is called when production deferrable_option is entered.
func (s *BaseDmSqlParserListener) EnterDeferrable_option(ctx *Deferrable_optionContext) {}

// ExitDeferrable_option is called when production deferrable_option is exited.
func (s *BaseDmSqlParserListener) ExitDeferrable_option(ctx *Deferrable_optionContext) {}

// EnterConstraint_check_time is called when production constraint_check_time is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_check_time(ctx *Constraint_check_timeContext) {}

// ExitConstraint_check_time is called when production constraint_check_time is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_check_time(ctx *Constraint_check_timeContext) {}

// EnterTable_constraint_clause is called when production table_constraint_clause is entered.
func (s *BaseDmSqlParserListener) EnterTable_constraint_clause(ctx *Table_constraint_clauseContext) {}

// ExitTable_constraint_clause is called when production table_constraint_clause is exited.
func (s *BaseDmSqlParserListener) ExitTable_constraint_clause(ctx *Table_constraint_clauseContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BaseDmSqlParserListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BaseDmSqlParserListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterUsing_index_clause is called when production using_index_clause is entered.
func (s *BaseDmSqlParserListener) EnterUsing_index_clause(ctx *Using_index_clauseContext) {}

// ExitUsing_index_clause is called when production using_index_clause is exited.
func (s *BaseDmSqlParserListener) ExitUsing_index_clause(ctx *Using_index_clauseContext) {}

// EnterForeign_key_clause is called when production foreign_key_clause is entered.
func (s *BaseDmSqlParserListener) EnterForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// ExitForeign_key_clause is called when production foreign_key_clause is exited.
func (s *BaseDmSqlParserListener) ExitForeign_key_clause(ctx *Foreign_key_clauseContext) {}

// EnterAlter_trigger_stmt is called when production alter_trigger_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) {}

// ExitAlter_trigger_stmt is called when production alter_trigger_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_trigger_stmt(ctx *Alter_trigger_stmtContext) {}

// EnterAlter_trigger_option is called when production alter_trigger_option is entered.
func (s *BaseDmSqlParserListener) EnterAlter_trigger_option(ctx *Alter_trigger_optionContext) {}

// ExitAlter_trigger_option is called when production alter_trigger_option is exited.
func (s *BaseDmSqlParserListener) ExitAlter_trigger_option(ctx *Alter_trigger_optionContext) {}

// EnterAlter_table_partition_action_options is called when production alter_table_partition_action_options is entered.
func (s *BaseDmSqlParserListener) EnterAlter_table_partition_action_options(ctx *Alter_table_partition_action_optionsContext) {
}

// ExitAlter_table_partition_action_options is called when production alter_table_partition_action_options is exited.
func (s *BaseDmSqlParserListener) ExitAlter_table_partition_action_options(ctx *Alter_table_partition_action_optionsContext) {
}

// EnterRefresh_materialized_view_stmt is called when production refresh_materialized_view_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) {
}

// ExitRefresh_materialized_view_stmt is called when production refresh_materialized_view_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRefresh_materialized_view_stmt(ctx *Refresh_materialized_view_stmtContext) {
}

// EnterComplete_del_null is called when production complete_del_null is entered.
func (s *BaseDmSqlParserListener) EnterComplete_del_null(ctx *Complete_del_nullContext) {}

// ExitComplete_del_null is called when production complete_del_null is exited.
func (s *BaseDmSqlParserListener) ExitComplete_del_null(ctx *Complete_del_nullContext) {}

// EnterRefresh_complete_del is called when production refresh_complete_del is entered.
func (s *BaseDmSqlParserListener) EnterRefresh_complete_del(ctx *Refresh_complete_delContext) {}

// ExitRefresh_complete_del is called when production refresh_complete_del is exited.
func (s *BaseDmSqlParserListener) ExitRefresh_complete_del(ctx *Refresh_complete_delContext) {}

// EnterAlter_materialized_view_stmt is called when production alter_materialized_view_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_materialized_view_stmt(ctx *Alter_materialized_view_stmtContext) {
}

// ExitAlter_materialized_view_stmt is called when production alter_materialized_view_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_materialized_view_stmt(ctx *Alter_materialized_view_stmtContext) {
}

// EnterAlter_view_stmt is called when production alter_view_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_view_stmt(ctx *Alter_view_stmtContext) {}

// ExitAlter_view_stmt is called when production alter_view_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_view_stmt(ctx *Alter_view_stmtContext) {}

// EnterAlter_view_action is called when production alter_view_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_view_action(ctx *Alter_view_actionContext) {}

// ExitAlter_view_action is called when production alter_view_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_view_action(ctx *Alter_view_actionContext) {}

// EnterCons_novalidate is called when production cons_novalidate is entered.
func (s *BaseDmSqlParserListener) EnterCons_novalidate(ctx *Cons_novalidateContext) {}

// ExitCons_novalidate is called when production cons_novalidate is exited.
func (s *BaseDmSqlParserListener) ExitCons_novalidate(ctx *Cons_novalidateContext) {}

// EnterView_constraint_clause is called when production view_constraint_clause is entered.
func (s *BaseDmSqlParserListener) EnterView_constraint_clause(ctx *View_constraint_clauseContext) {}

// ExitView_constraint_clause is called when production view_constraint_clause is exited.
func (s *BaseDmSqlParserListener) ExitView_constraint_clause(ctx *View_constraint_clauseContext) {}

// EnterView_constraint is called when production view_constraint is entered.
func (s *BaseDmSqlParserListener) EnterView_constraint(ctx *View_constraintContext) {}

// ExitView_constraint is called when production view_constraint is exited.
func (s *BaseDmSqlParserListener) ExitView_constraint(ctx *View_constraintContext) {}

// EnterView_unique_spec is called when production view_unique_spec is entered.
func (s *BaseDmSqlParserListener) EnterView_unique_spec(ctx *View_unique_specContext) {}

// ExitView_unique_spec is called when production view_unique_spec is exited.
func (s *BaseDmSqlParserListener) ExitView_unique_spec(ctx *View_unique_specContext) {}

// EnterView_refs_spec is called when production view_refs_spec is entered.
func (s *BaseDmSqlParserListener) EnterView_refs_spec(ctx *View_refs_specContext) {}

// ExitView_refs_spec is called when production view_refs_spec is exited.
func (s *BaseDmSqlParserListener) ExitView_refs_spec(ctx *View_refs_specContext) {}

// EnterView_refs_spec_action is called when production view_refs_spec_action is entered.
func (s *BaseDmSqlParserListener) EnterView_refs_spec_action(ctx *View_refs_spec_actionContext) {}

// ExitView_refs_spec_action is called when production view_refs_spec_action is exited.
func (s *BaseDmSqlParserListener) ExitView_refs_spec_action(ctx *View_refs_spec_actionContext) {}

// EnterCall_proc_stmt is called when production call_proc_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCall_proc_stmt(ctx *Call_proc_stmtContext) {}

// ExitCall_proc_stmt is called when production call_proc_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCall_proc_stmt(ctx *Call_proc_stmtContext) {}

// EnterRaw_call_proc_stmt is called when production raw_call_proc_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRaw_call_proc_stmt(ctx *Raw_call_proc_stmtContext) {}

// ExitRaw_call_proc_stmt is called when production raw_call_proc_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRaw_call_proc_stmt(ctx *Raw_call_proc_stmtContext) {}

// EnterCall_proc_stmt_2 is called when production call_proc_stmt_2 is entered.
func (s *BaseDmSqlParserListener) EnterCall_proc_stmt_2(ctx *Call_proc_stmt_2Context) {}

// ExitCall_proc_stmt_2 is called when production call_proc_stmt_2 is exited.
func (s *BaseDmSqlParserListener) ExitCall_proc_stmt_2(ctx *Call_proc_stmt_2Context) {}

// EnterExec_proc_stmt is called when production exec_proc_stmt is entered.
func (s *BaseDmSqlParserListener) EnterExec_proc_stmt(ctx *Exec_proc_stmtContext) {}

// ExitExec_proc_stmt is called when production exec_proc_stmt is exited.
func (s *BaseDmSqlParserListener) ExitExec_proc_stmt(ctx *Exec_proc_stmtContext) {}

// EnterDblink_clause is called when production dblink_clause is entered.
func (s *BaseDmSqlParserListener) EnterDblink_clause(ctx *Dblink_clauseContext) {}

// ExitDblink_clause is called when production dblink_clause is exited.
func (s *BaseDmSqlParserListener) ExitDblink_clause(ctx *Dblink_clauseContext) {}

// EnterDblink_clause2 is called when production dblink_clause2 is entered.
func (s *BaseDmSqlParserListener) EnterDblink_clause2(ctx *Dblink_clause2Context) {}

// ExitDblink_clause2 is called when production dblink_clause2 is exited.
func (s *BaseDmSqlParserListener) ExitDblink_clause2(ctx *Dblink_clause2Context) {}

// EnterParam_list is called when production param_list is entered.
func (s *BaseDmSqlParserListener) EnterParam_list(ctx *Param_listContext) {}

// ExitParam_list is called when production param_list is exited.
func (s *BaseDmSqlParserListener) ExitParam_list(ctx *Param_listContext) {}

// EnterParam is called when production param is entered.
func (s *BaseDmSqlParserListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseDmSqlParserListener) ExitParam(ctx *ParamContext) {}

// EnterRaw_exp_list is called when production raw_exp_list is entered.
func (s *BaseDmSqlParserListener) EnterRaw_exp_list(ctx *Raw_exp_listContext) {}

// ExitRaw_exp_list is called when production raw_exp_list is exited.
func (s *BaseDmSqlParserListener) ExitRaw_exp_list(ctx *Raw_exp_listContext) {}

// EnterExp_list_2 is called when production exp_list_2 is entered.
func (s *BaseDmSqlParserListener) EnterExp_list_2(ctx *Exp_list_2Context) {}

// ExitExp_list_2 is called when production exp_list_2 is exited.
func (s *BaseDmSqlParserListener) ExitExp_list_2(ctx *Exp_list_2Context) {}

// EnterExp_list is called when production exp_list is entered.
func (s *BaseDmSqlParserListener) EnterExp_list(ctx *Exp_listContext) {}

// ExitExp_list is called when production exp_list is exited.
func (s *BaseDmSqlParserListener) ExitExp_list(ctx *Exp_listContext) {}

// EnterIns_exp_list is called when production ins_exp_list is entered.
func (s *BaseDmSqlParserListener) EnterIns_exp_list(ctx *Ins_exp_listContext) {}

// ExitIns_exp_list is called when production ins_exp_list is exited.
func (s *BaseDmSqlParserListener) ExitIns_exp_list(ctx *Ins_exp_listContext) {}

// EnterLt_exp is called when production lt_exp is entered.
func (s *BaseDmSqlParserListener) EnterLt_exp(ctx *Lt_expContext) {}

// ExitLt_exp is called when production lt_exp is exited.
func (s *BaseDmSqlParserListener) ExitLt_exp(ctx *Lt_expContext) {}

// EnterRange_partition_exp is called when production range_partition_exp is entered.
func (s *BaseDmSqlParserListener) EnterRange_partition_exp(ctx *Range_partition_expContext) {}

// ExitRange_partition_exp is called when production range_partition_exp is exited.
func (s *BaseDmSqlParserListener) ExitRange_partition_exp(ctx *Range_partition_expContext) {}

// EnterRange_partition_exp_list is called when production range_partition_exp_list is entered.
func (s *BaseDmSqlParserListener) EnterRange_partition_exp_list(ctx *Range_partition_exp_listContext) {
}

// ExitRange_partition_exp_list is called when production range_partition_exp_list is exited.
func (s *BaseDmSqlParserListener) ExitRange_partition_exp_list(ctx *Range_partition_exp_listContext) {
}

// EnterList_partition_exp is called when production list_partition_exp is entered.
func (s *BaseDmSqlParserListener) EnterList_partition_exp(ctx *List_partition_expContext) {}

// ExitList_partition_exp is called when production list_partition_exp is exited.
func (s *BaseDmSqlParserListener) ExitList_partition_exp(ctx *List_partition_expContext) {}

// EnterList_partition_exp_list is called when production list_partition_exp_list is entered.
func (s *BaseDmSqlParserListener) EnterList_partition_exp_list(ctx *List_partition_exp_listContext) {}

// ExitList_partition_exp_list is called when production list_partition_exp_list is exited.
func (s *BaseDmSqlParserListener) ExitList_partition_exp_list(ctx *List_partition_exp_listContext) {}

// EnterList_partition_value_list is called when production list_partition_value_list is entered.
func (s *BaseDmSqlParserListener) EnterList_partition_value_list(ctx *List_partition_value_listContext) {
}

// ExitList_partition_value_list is called when production list_partition_value_list is exited.
func (s *BaseDmSqlParserListener) ExitList_partition_value_list(ctx *List_partition_value_listContext) {
}

// EnterClose_cursor_stmt is called when production close_cursor_stmt is entered.
func (s *BaseDmSqlParserListener) EnterClose_cursor_stmt(ctx *Close_cursor_stmtContext) {}

// ExitClose_cursor_stmt is called when production close_cursor_stmt is exited.
func (s *BaseDmSqlParserListener) ExitClose_cursor_stmt(ctx *Close_cursor_stmtContext) {}

// EnterClose_cursor_statement is called when production close_cursor_statement is entered.
func (s *BaseDmSqlParserListener) EnterClose_cursor_statement(ctx *Close_cursor_statementContext) {}

// ExitClose_cursor_statement is called when production close_cursor_statement is exited.
func (s *BaseDmSqlParserListener) ExitClose_cursor_statement(ctx *Close_cursor_statementContext) {}

// EnterBegin_trans_stmt is called when production begin_trans_stmt is entered.
func (s *BaseDmSqlParserListener) EnterBegin_trans_stmt(ctx *Begin_trans_stmtContext) {}

// ExitBegin_trans_stmt is called when production begin_trans_stmt is exited.
func (s *BaseDmSqlParserListener) ExitBegin_trans_stmt(ctx *Begin_trans_stmtContext) {}

// EnterCommit_trans_stmt is called when production commit_trans_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCommit_trans_stmt(ctx *Commit_trans_stmtContext) {}

// ExitCommit_trans_stmt is called when production commit_trans_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCommit_trans_stmt(ctx *Commit_trans_stmtContext) {}

// EnterCommit_head is called when production commit_head is entered.
func (s *BaseDmSqlParserListener) EnterCommit_head(ctx *Commit_headContext) {}

// ExitCommit_head is called when production commit_head is exited.
func (s *BaseDmSqlParserListener) ExitCommit_head(ctx *Commit_headContext) {}

// EnterCommit_tail is called when production commit_tail is entered.
func (s *BaseDmSqlParserListener) EnterCommit_tail(ctx *Commit_tailContext) {}

// ExitCommit_tail is called when production commit_tail is exited.
func (s *BaseDmSqlParserListener) ExitCommit_tail(ctx *Commit_tailContext) {}

// EnterCommit_wait_immed_option is called when production commit_wait_immed_option is entered.
func (s *BaseDmSqlParserListener) EnterCommit_wait_immed_option(ctx *Commit_wait_immed_optionContext) {
}

// ExitCommit_wait_immed_option is called when production commit_wait_immed_option is exited.
func (s *BaseDmSqlParserListener) ExitCommit_wait_immed_option(ctx *Commit_wait_immed_optionContext) {
}

// EnterConnect_stmt is called when production connect_stmt is entered.
func (s *BaseDmSqlParserListener) EnterConnect_stmt(ctx *Connect_stmtContext) {}

// ExitConnect_stmt is called when production connect_stmt is exited.
func (s *BaseDmSqlParserListener) ExitConnect_stmt(ctx *Connect_stmtContext) {}

// EnterPassword is called when production password is entered.
func (s *BaseDmSqlParserListener) EnterPassword(ctx *PasswordContext) {}

// ExitPassword is called when production password is exited.
func (s *BaseDmSqlParserListener) ExitPassword(ctx *PasswordContext) {}

// EnterTs_storage is called when production ts_storage is entered.
func (s *BaseDmSqlParserListener) EnterTs_storage(ctx *Ts_storageContext) {}

// ExitTs_storage is called when production ts_storage is exited.
func (s *BaseDmSqlParserListener) ExitTs_storage(ctx *Ts_storageContext) {}

// EnterTs_storage_clause is called when production ts_storage_clause is entered.
func (s *BaseDmSqlParserListener) EnterTs_storage_clause(ctx *Ts_storage_clauseContext) {}

// ExitTs_storage_clause is called when production ts_storage_clause is exited.
func (s *BaseDmSqlParserListener) ExitTs_storage_clause(ctx *Ts_storage_clauseContext) {}

// EnterCreate_tablespace_stmt is called when production create_tablespace_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) {}

// ExitCreate_tablespace_stmt is called when production create_tablespace_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_tablespace_stmt(ctx *Create_tablespace_stmtContext) {}

// EnterCtss_with_clause is called when production ctss_with_clause is entered.
func (s *BaseDmSqlParserListener) EnterCtss_with_clause(ctx *Ctss_with_clauseContext) {}

// ExitCtss_with_clause is called when production ctss_with_clause is exited.
func (s *BaseDmSqlParserListener) ExitCtss_with_clause(ctx *Ctss_with_clauseContext) {}

// EnterCreate_tablespace_set_stmt is called when production create_tablespace_set_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_tablespace_set_stmt(ctx *Create_tablespace_set_stmtContext) {
}

// ExitCreate_tablespace_set_stmt is called when production create_tablespace_set_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_tablespace_set_stmt(ctx *Create_tablespace_set_stmtContext) {
}

// EnterAlter_tablespace_set_stmt is called when production alter_tablespace_set_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_tablespace_set_stmt(ctx *Alter_tablespace_set_stmtContext) {
}

// ExitAlter_tablespace_set_stmt is called when production alter_tablespace_set_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_tablespace_set_stmt(ctx *Alter_tablespace_set_stmtContext) {
}

// EnterCache is called when production cache is entered.
func (s *BaseDmSqlParserListener) EnterCache(ctx *CacheContext) {}

// ExitCache is called when production cache is exited.
func (s *BaseDmSqlParserListener) ExitCache(ctx *CacheContext) {}

// EnterAlter_tablespace_stmt is called when production alter_tablespace_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) {}

// ExitAlter_tablespace_stmt is called when production alter_tablespace_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_tablespace_stmt(ctx *Alter_tablespace_stmtContext) {}

// EnterKeep is called when production keep is entered.
func (s *BaseDmSqlParserListener) EnterKeep(ctx *KeepContext) {}

// ExitKeep is called when production keep is exited.
func (s *BaseDmSqlParserListener) ExitKeep(ctx *KeepContext) {}

// EnterAlter_tablespace_action is called when production alter_tablespace_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_tablespace_action(ctx *Alter_tablespace_actionContext) {}

// ExitAlter_tablespace_action is called when production alter_tablespace_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_tablespace_action(ctx *Alter_tablespace_actionContext) {}

// EnterFile_list is called when production file_list is entered.
func (s *BaseDmSqlParserListener) EnterFile_list(ctx *File_listContext) {}

// ExitFile_list is called when production file_list is exited.
func (s *BaseDmSqlParserListener) ExitFile_list(ctx *File_listContext) {}

// EnterPathname_list is called when production pathname_list is entered.
func (s *BaseDmSqlParserListener) EnterPathname_list(ctx *Pathname_listContext) {}

// ExitPathname_list is called when production pathname_list is exited.
func (s *BaseDmSqlParserListener) ExitPathname_list(ctx *Pathname_listContext) {}

// EnterInteger_list is called when production integer_list is entered.
func (s *BaseDmSqlParserListener) EnterInteger_list(ctx *Integer_listContext) {}

// ExitInteger_list is called when production integer_list is exited.
func (s *BaseDmSqlParserListener) ExitInteger_list(ctx *Integer_listContext) {}

// EnterFile is called when production file is entered.
func (s *BaseDmSqlParserListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseDmSqlParserListener) ExitFile(ctx *FileContext) {}

// EnterMirror is called when production mirror is entered.
func (s *BaseDmSqlParserListener) EnterMirror(ctx *MirrorContext) {}

// ExitMirror is called when production mirror is exited.
func (s *BaseDmSqlParserListener) ExitMirror(ctx *MirrorContext) {}

// EnterAutoextend_nextsize is called when production autoextend_nextsize is entered.
func (s *BaseDmSqlParserListener) EnterAutoextend_nextsize(ctx *Autoextend_nextsizeContext) {}

// ExitAutoextend_nextsize is called when production autoextend_nextsize is exited.
func (s *BaseDmSqlParserListener) ExitAutoextend_nextsize(ctx *Autoextend_nextsizeContext) {}

// EnterAutoextend_maxsize is called when production autoextend_maxsize is entered.
func (s *BaseDmSqlParserListener) EnterAutoextend_maxsize(ctx *Autoextend_maxsizeContext) {}

// ExitAutoextend_maxsize is called when production autoextend_maxsize is exited.
func (s *BaseDmSqlParserListener) ExitAutoextend_maxsize(ctx *Autoextend_maxsizeContext) {}

// EnterAutoextend is called when production autoextend is entered.
func (s *BaseDmSqlParserListener) EnterAutoextend(ctx *AutoextendContext) {}

// ExitAutoextend is called when production autoextend is exited.
func (s *BaseDmSqlParserListener) ExitAutoextend(ctx *AutoextendContext) {}

// EnterOn_raft is called when production on_raft is entered.
func (s *BaseDmSqlParserListener) EnterOn_raft(ctx *On_raftContext) {}

// ExitOn_raft is called when production on_raft is exited.
func (s *BaseDmSqlParserListener) ExitOn_raft(ctx *On_raftContext) {}

// EnterArchfile is called when production archfile is entered.
func (s *BaseDmSqlParserListener) EnterArchfile(ctx *ArchfileContext) {}

// ExitArchfile is called when production archfile is exited.
func (s *BaseDmSqlParserListener) ExitArchfile(ctx *ArchfileContext) {}

// EnterArchflag is called when production archflag is entered.
func (s *BaseDmSqlParserListener) EnterArchflag(ctx *ArchflagContext) {}

// ExitArchflag is called when production archflag is exited.
func (s *BaseDmSqlParserListener) ExitArchflag(ctx *ArchflagContext) {}

// EnterArchstyle_options is called when production archstyle_options is entered.
func (s *BaseDmSqlParserListener) EnterArchstyle_options(ctx *Archstyle_optionsContext) {}

// ExitArchstyle_options is called when production archstyle_options is exited.
func (s *BaseDmSqlParserListener) ExitArchstyle_options(ctx *Archstyle_optionsContext) {}

// EnterArchstyle is called when production archstyle is entered.
func (s *BaseDmSqlParserListener) EnterArchstyle(ctx *ArchstyleContext) {}

// ExitArchstyle is called when production archstyle is exited.
func (s *BaseDmSqlParserListener) ExitArchstyle(ctx *ArchstyleContext) {}

// EnterArchdir is called when production archdir is entered.
func (s *BaseDmSqlParserListener) EnterArchdir(ctx *ArchdirContext) {}

// ExitArchdir is called when production archdir is exited.
func (s *BaseDmSqlParserListener) ExitArchdir(ctx *ArchdirContext) {}

// EnterBakfile is called when production bakfile is entered.
func (s *BaseDmSqlParserListener) EnterBakfile(ctx *BakfileContext) {}

// ExitBakfile is called when production bakfile is exited.
func (s *BaseDmSqlParserListener) ExitBakfile(ctx *BakfileContext) {}

// EnterParameters_option_list is called when production parameters_option_list is entered.
func (s *BaseDmSqlParserListener) EnterParameters_option_list(ctx *Parameters_option_listContext) {}

// ExitParameters_option_list is called when production parameters_option_list is exited.
func (s *BaseDmSqlParserListener) ExitParameters_option_list(ctx *Parameters_option_listContext) {}

// EnterParameter_option_list is called when production parameter_option_list is entered.
func (s *BaseDmSqlParserListener) EnterParameter_option_list(ctx *Parameter_option_listContext) {}

// ExitParameter_option_list is called when production parameter_option_list is exited.
func (s *BaseDmSqlParserListener) ExitParameter_option_list(ctx *Parameter_option_listContext) {}

// EnterParameter_option is called when production parameter_option is entered.
func (s *BaseDmSqlParserListener) EnterParameter_option(ctx *Parameter_optionContext) {}

// ExitParameter_option is called when production parameter_option is exited.
func (s *BaseDmSqlParserListener) ExitParameter_option(ctx *Parameter_optionContext) {}

// EnterPathname is called when production pathname is entered.
func (s *BaseDmSqlParserListener) EnterPathname(ctx *PathnameContext) {}

// ExitPathname is called when production pathname is exited.
func (s *BaseDmSqlParserListener) ExitPathname(ctx *PathnameContext) {}

// EnterPathname_options is called when production pathname_options is entered.
func (s *BaseDmSqlParserListener) EnterPathname_options(ctx *Pathname_optionsContext) {}

// ExitPathname_options is called when production pathname_options is exited.
func (s *BaseDmSqlParserListener) ExitPathname_options(ctx *Pathname_optionsContext) {}

// EnterBackup_stmt is called when production backup_stmt is entered.
func (s *BaseDmSqlParserListener) EnterBackup_stmt(ctx *Backup_stmtContext) {}

// ExitBackup_stmt is called when production backup_stmt is exited.
func (s *BaseDmSqlParserListener) ExitBackup_stmt(ctx *Backup_stmtContext) {}

// EnterBack_range_option is called when production back_range_option is entered.
func (s *BaseDmSqlParserListener) EnterBack_range_option(ctx *Back_range_optionContext) {}

// ExitBack_range_option is called when production back_range_option is exited.
func (s *BaseDmSqlParserListener) ExitBack_range_option(ctx *Back_range_optionContext) {}

// EnterBack_archive_spec_null is called when production back_archive_spec_null is entered.
func (s *BaseDmSqlParserListener) EnterBack_archive_spec_null(ctx *Back_archive_spec_nullContext) {}

// ExitBack_archive_spec_null is called when production back_archive_spec_null is exited.
func (s *BaseDmSqlParserListener) ExitBack_archive_spec_null(ctx *Back_archive_spec_nullContext) {}

// EnterNot_backed_up is called when production not_backed_up is entered.
func (s *BaseDmSqlParserListener) EnterNot_backed_up(ctx *Not_backed_upContext) {}

// ExitNot_backed_up is called when production not_backed_up is exited.
func (s *BaseDmSqlParserListener) ExitNot_backed_up(ctx *Not_backed_upContext) {}

// EnterArchive_spec is called when production archive_spec is entered.
func (s *BaseDmSqlParserListener) EnterArchive_spec(ctx *Archive_specContext) {}

// ExitArchive_spec is called when production archive_spec is exited.
func (s *BaseDmSqlParserListener) ExitArchive_spec(ctx *Archive_specContext) {}

// EnterSpec_lsn is called when production spec_lsn is entered.
func (s *BaseDmSqlParserListener) EnterSpec_lsn(ctx *Spec_lsnContext) {}

// ExitSpec_lsn is called when production spec_lsn is exited.
func (s *BaseDmSqlParserListener) ExitSpec_lsn(ctx *Spec_lsnContext) {}

// EnterBackup_delete_archive is called when production backup_delete_archive is entered.
func (s *BaseDmSqlParserListener) EnterBackup_delete_archive(ctx *Backup_delete_archiveContext) {}

// ExitBackup_delete_archive is called when production backup_delete_archive is exited.
func (s *BaseDmSqlParserListener) ExitBackup_delete_archive(ctx *Backup_delete_archiveContext) {}

// EnterBackup_options is called when production backup_options is entered.
func (s *BaseDmSqlParserListener) EnterBackup_options(ctx *Backup_optionsContext) {}

// ExitBackup_options is called when production backup_options is exited.
func (s *BaseDmSqlParserListener) ExitBackup_options(ctx *Backup_optionsContext) {}

// EnterCumulative is called when production cumulative is entered.
func (s *BaseDmSqlParserListener) EnterCumulative(ctx *CumulativeContext) {}

// ExitCumulative is called when production cumulative is exited.
func (s *BaseDmSqlParserListener) ExitCumulative(ctx *CumulativeContext) {}

// EnterWith_bak_dir_list is called when production with_bak_dir_list is entered.
func (s *BaseDmSqlParserListener) EnterWith_bak_dir_list(ctx *With_bak_dir_listContext) {}

// ExitWith_bak_dir_list is called when production with_bak_dir_list is exited.
func (s *BaseDmSqlParserListener) ExitWith_bak_dir_list(ctx *With_bak_dir_listContext) {}

// EnterBase_on_backup is called when production base_on_backup is entered.
func (s *BaseDmSqlParserListener) EnterBase_on_backup(ctx *Base_on_backupContext) {}

// ExitBase_on_backup is called when production base_on_backup is exited.
func (s *BaseDmSqlParserListener) ExitBase_on_backup(ctx *Base_on_backupContext) {}

// EnterBackup_to_options is called when production backup_to_options is entered.
func (s *BaseDmSqlParserListener) EnterBackup_to_options(ctx *Backup_to_optionsContext) {}

// ExitBackup_to_options is called when production backup_to_options is exited.
func (s *BaseDmSqlParserListener) ExitBackup_to_options(ctx *Backup_to_optionsContext) {}

// EnterBackup_path_null is called when production backup_path_null is entered.
func (s *BaseDmSqlParserListener) EnterBackup_path_null(ctx *Backup_path_nullContext) {}

// ExitBackup_path_null is called when production backup_path_null is exited.
func (s *BaseDmSqlParserListener) ExitBackup_path_null(ctx *Backup_path_nullContext) {}

// EnterDevice_type is called when production device_type is entered.
func (s *BaseDmSqlParserListener) EnterDevice_type(ctx *Device_typeContext) {}

// ExitDevice_type is called when production device_type is exited.
func (s *BaseDmSqlParserListener) ExitDevice_type(ctx *Device_typeContext) {}

// EnterParms_command is called when production parms_command is entered.
func (s *BaseDmSqlParserListener) EnterParms_command(ctx *Parms_commandContext) {}

// ExitParms_command is called when production parms_command is exited.
func (s *BaseDmSqlParserListener) ExitParms_command(ctx *Parms_commandContext) {}

// EnterMedia_name is called when production media_name is entered.
func (s *BaseDmSqlParserListener) EnterMedia_name(ctx *Media_nameContext) {}

// ExitMedia_name is called when production media_name is exited.
func (s *BaseDmSqlParserListener) ExitMedia_name(ctx *Media_nameContext) {}

// EnterBackup_desc_options is called when production backup_desc_options is entered.
func (s *BaseDmSqlParserListener) EnterBackup_desc_options(ctx *Backup_desc_optionsContext) {}

// ExitBackup_desc_options is called when production backup_desc_options is exited.
func (s *BaseDmSqlParserListener) ExitBackup_desc_options(ctx *Backup_desc_optionsContext) {}

// EnterBackup_desc is called when production backup_desc is entered.
func (s *BaseDmSqlParserListener) EnterBackup_desc(ctx *Backup_descContext) {}

// ExitBackup_desc is called when production backup_desc is exited.
func (s *BaseDmSqlParserListener) ExitBackup_desc(ctx *Backup_descContext) {}

// EnterBackup_maxsize is called when production backup_maxsize is entered.
func (s *BaseDmSqlParserListener) EnterBackup_maxsize(ctx *Backup_maxsizeContext) {}

// ExitBackup_maxsize is called when production backup_maxsize is exited.
func (s *BaseDmSqlParserListener) ExitBackup_maxsize(ctx *Backup_maxsizeContext) {}

// EnterBackup_limit is called when production backup_limit is entered.
func (s *BaseDmSqlParserListener) EnterBackup_limit(ctx *Backup_limitContext) {}

// ExitBackup_limit is called when production backup_limit is exited.
func (s *BaseDmSqlParserListener) ExitBackup_limit(ctx *Backup_limitContext) {}

// EnterBackup_identified is called when production backup_identified is entered.
func (s *BaseDmSqlParserListener) EnterBackup_identified(ctx *Backup_identifiedContext) {}

// ExitBackup_identified is called when production backup_identified is exited.
func (s *BaseDmSqlParserListener) ExitBackup_identified(ctx *Backup_identifiedContext) {}

// EnterBackup_compressed is called when production backup_compressed is entered.
func (s *BaseDmSqlParserListener) EnterBackup_compressed(ctx *Backup_compressedContext) {}

// ExitBackup_compressed is called when production backup_compressed is exited.
func (s *BaseDmSqlParserListener) ExitBackup_compressed(ctx *Backup_compressedContext) {}

// EnterBackup_without is called when production backup_without is entered.
func (s *BaseDmSqlParserListener) EnterBackup_without(ctx *Backup_withoutContext) {}

// ExitBackup_without is called when production backup_without is exited.
func (s *BaseDmSqlParserListener) ExitBackup_without(ctx *Backup_withoutContext) {}

// EnterBackup_tsk_thread_num_null is called when production backup_tsk_thread_num_null is entered.
func (s *BaseDmSqlParserListener) EnterBackup_tsk_thread_num_null(ctx *Backup_tsk_thread_num_nullContext) {
}

// ExitBackup_tsk_thread_num_null is called when production backup_tsk_thread_num_null is exited.
func (s *BaseDmSqlParserListener) ExitBackup_tsk_thread_num_null(ctx *Backup_tsk_thread_num_nullContext) {
}

// EnterBackup_parallel_dir is called when production backup_parallel_dir is entered.
func (s *BaseDmSqlParserListener) EnterBackup_parallel_dir(ctx *Backup_parallel_dirContext) {}

// ExitBackup_parallel_dir is called when production backup_parallel_dir is exited.
func (s *BaseDmSqlParserListener) ExitBackup_parallel_dir(ctx *Backup_parallel_dirContext) {}

// EnterBackup_trace_file_level is called when production backup_trace_file_level is entered.
func (s *BaseDmSqlParserListener) EnterBackup_trace_file_level(ctx *Backup_trace_file_levelContext) {}

// ExitBackup_trace_file_level is called when production backup_trace_file_level is exited.
func (s *BaseDmSqlParserListener) ExitBackup_trace_file_level(ctx *Backup_trace_file_levelContext) {}

// EnterRestore_stmt is called when production restore_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRestore_stmt(ctx *Restore_stmtContext) {}

// ExitRestore_stmt is called when production restore_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRestore_stmt(ctx *Restore_stmtContext) {}

// EnterRestore_datafile_lst is called when production restore_datafile_lst is entered.
func (s *BaseDmSqlParserListener) EnterRestore_datafile_lst(ctx *Restore_datafile_lstContext) {}

// ExitRestore_datafile_lst is called when production restore_datafile_lst is exited.
func (s *BaseDmSqlParserListener) ExitRestore_datafile_lst(ctx *Restore_datafile_lstContext) {}

// EnterRestore_mapped_file is called when production restore_mapped_file is entered.
func (s *BaseDmSqlParserListener) EnterRestore_mapped_file(ctx *Restore_mapped_fileContext) {}

// ExitRestore_mapped_file is called when production restore_mapped_file is exited.
func (s *BaseDmSqlParserListener) ExitRestore_mapped_file(ctx *Restore_mapped_fileContext) {}

// EnterMapped_file is called when production mapped_file is entered.
func (s *BaseDmSqlParserListener) EnterMapped_file(ctx *Mapped_fileContext) {}

// ExitMapped_file is called when production mapped_file is exited.
func (s *BaseDmSqlParserListener) ExitMapped_file(ctx *Mapped_fileContext) {}

// EnterRes_struct is called when production res_struct is entered.
func (s *BaseDmSqlParserListener) EnterRes_struct(ctx *Res_structContext) {}

// ExitRes_struct is called when production res_struct is exited.
func (s *BaseDmSqlParserListener) ExitRes_struct(ctx *Res_structContext) {}

// EnterTsk_thread_num is called when production tsk_thread_num is entered.
func (s *BaseDmSqlParserListener) EnterTsk_thread_num(ctx *Tsk_thread_numContext) {}

// ExitTsk_thread_num is called when production tsk_thread_num is exited.
func (s *BaseDmSqlParserListener) ExitTsk_thread_num(ctx *Tsk_thread_numContext) {}

// EnterRestore_tsk_thread_num_null is called when production restore_tsk_thread_num_null is entered.
func (s *BaseDmSqlParserListener) EnterRestore_tsk_thread_num_null(ctx *Restore_tsk_thread_num_nullContext) {
}

// ExitRestore_tsk_thread_num_null is called when production restore_tsk_thread_num_null is exited.
func (s *BaseDmSqlParserListener) ExitRestore_tsk_thread_num_null(ctx *Restore_tsk_thread_num_nullContext) {
}

// EnterRestore_parallel is called when production restore_parallel is entered.
func (s *BaseDmSqlParserListener) EnterRestore_parallel(ctx *Restore_parallelContext) {}

// ExitRestore_parallel is called when production restore_parallel is exited.
func (s *BaseDmSqlParserListener) ExitRestore_parallel(ctx *Restore_parallelContext) {}

// EnterFull_table_name_options is called when production full_table_name_options is entered.
func (s *BaseDmSqlParserListener) EnterFull_table_name_options(ctx *Full_table_name_optionsContext) {}

// ExitFull_table_name_options is called when production full_table_name_options is exited.
func (s *BaseDmSqlParserListener) ExitFull_table_name_options(ctx *Full_table_name_optionsContext) {}

// EnterRes_without_index_constraint is called when production res_without_index_constraint is entered.
func (s *BaseDmSqlParserListener) EnterRes_without_index_constraint(ctx *Res_without_index_constraintContext) {
}

// ExitRes_without_index_constraint is called when production res_without_index_constraint is exited.
func (s *BaseDmSqlParserListener) ExitRes_without_index_constraint(ctx *Res_without_index_constraintContext) {
}

// EnterRestore_from is called when production restore_from is entered.
func (s *BaseDmSqlParserListener) EnterRestore_from(ctx *Restore_fromContext) {}

// ExitRestore_from is called when production restore_from is exited.
func (s *BaseDmSqlParserListener) ExitRestore_from(ctx *Restore_fromContext) {}

// EnterRes_until is called when production res_until is entered.
func (s *BaseDmSqlParserListener) EnterRes_until(ctx *Res_untilContext) {}

// ExitRes_until is called when production res_until is exited.
func (s *BaseDmSqlParserListener) ExitRes_until(ctx *Res_untilContext) {}

// EnterRestore_file_list_options is called when production restore_file_list_options is entered.
func (s *BaseDmSqlParserListener) EnterRestore_file_list_options(ctx *Restore_file_list_optionsContext) {
}

// ExitRestore_file_list_options is called when production restore_file_list_options is exited.
func (s *BaseDmSqlParserListener) ExitRestore_file_list_options(ctx *Restore_file_list_optionsContext) {
}

// EnterMirror_file_list_options is called when production mirror_file_list_options is entered.
func (s *BaseDmSqlParserListener) EnterMirror_file_list_options(ctx *Mirror_file_list_optionsContext) {
}

// ExitMirror_file_list_options is called when production mirror_file_list_options is exited.
func (s *BaseDmSqlParserListener) ExitMirror_file_list_options(ctx *Mirror_file_list_optionsContext) {
}

// EnterRestore_trace_file_level is called when production restore_trace_file_level is entered.
func (s *BaseDmSqlParserListener) EnterRestore_trace_file_level(ctx *Restore_trace_file_levelContext) {
}

// ExitRestore_trace_file_level is called when production restore_trace_file_level is exited.
func (s *BaseDmSqlParserListener) ExitRestore_trace_file_level(ctx *Restore_trace_file_levelContext) {
}

// EnterRestore_file_list is called when production restore_file_list is entered.
func (s *BaseDmSqlParserListener) EnterRestore_file_list(ctx *Restore_file_listContext) {}

// ExitRestore_file_list is called when production restore_file_list is exited.
func (s *BaseDmSqlParserListener) ExitRestore_file_list(ctx *Restore_file_listContext) {}

// EnterRestore_file is called when production restore_file is entered.
func (s *BaseDmSqlParserListener) EnterRestore_file(ctx *Restore_fileContext) {}

// ExitRestore_file is called when production restore_file is exited.
func (s *BaseDmSqlParserListener) ExitRestore_file(ctx *Restore_fileContext) {}

// EnterMirror_file_list is called when production mirror_file_list is entered.
func (s *BaseDmSqlParserListener) EnterMirror_file_list(ctx *Mirror_file_listContext) {}

// ExitMirror_file_list is called when production mirror_file_list is exited.
func (s *BaseDmSqlParserListener) ExitMirror_file_list(ctx *Mirror_file_listContext) {}

// EnterMirror_file is called when production mirror_file is entered.
func (s *BaseDmSqlParserListener) EnterMirror_file(ctx *Mirror_fileContext) {}

// ExitMirror_file is called when production mirror_file is exited.
func (s *BaseDmSqlParserListener) ExitMirror_file(ctx *Mirror_fileContext) {}

// EnterWith_bak_arch_dir_list is called when production with_bak_arch_dir_list is entered.
func (s *BaseDmSqlParserListener) EnterWith_bak_arch_dir_list(ctx *With_bak_arch_dir_listContext) {}

// ExitWith_bak_arch_dir_list is called when production with_bak_arch_dir_list is exited.
func (s *BaseDmSqlParserListener) ExitWith_bak_arch_dir_list(ctx *With_bak_arch_dir_listContext) {}

// EnterRestore_identified is called when production restore_identified is entered.
func (s *BaseDmSqlParserListener) EnterRestore_identified(ctx *Restore_identifiedContext) {}

// ExitRestore_identified is called when production restore_identified is exited.
func (s *BaseDmSqlParserListener) ExitRestore_identified(ctx *Restore_identifiedContext) {}

// EnterCreate_func_stmt is called when production create_func_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_func_stmt(ctx *Create_func_stmtContext) {}

// ExitCreate_func_stmt is called when production create_func_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_func_stmt(ctx *Create_func_stmtContext) {}

// EnterFunc_aggr_clause is called when production func_aggr_clause is entered.
func (s *BaseDmSqlParserListener) EnterFunc_aggr_clause(ctx *Func_aggr_clauseContext) {}

// ExitFunc_aggr_clause is called when production func_aggr_clause is exited.
func (s *BaseDmSqlParserListener) ExitFunc_aggr_clause(ctx *Func_aggr_clauseContext) {}

// EnterPipelined_options is called when production pipelined_options is entered.
func (s *BaseDmSqlParserListener) EnterPipelined_options(ctx *Pipelined_optionsContext) {}

// ExitPipelined_options is called when production pipelined_options is exited.
func (s *BaseDmSqlParserListener) ExitPipelined_options(ctx *Pipelined_optionsContext) {}

// EnterReplace_option is called when production replace_option is entered.
func (s *BaseDmSqlParserListener) EnterReplace_option(ctx *Replace_optionContext) {}

// ExitReplace_option is called when production replace_option is exited.
func (s *BaseDmSqlParserListener) ExitReplace_option(ctx *Replace_optionContext) {}

// EnterEdit_options is called when production edit_options is entered.
func (s *BaseDmSqlParserListener) EnterEdit_options(ctx *Edit_optionsContext) {}

// ExitEdit_options is called when production edit_options is exited.
func (s *BaseDmSqlParserListener) ExitEdit_options(ctx *Edit_optionsContext) {}

// EnterEncryption_option is called when production encryption_option is entered.
func (s *BaseDmSqlParserListener) EnterEncryption_option(ctx *Encryption_optionContext) {}

// ExitEncryption_option is called when production encryption_option is exited.
func (s *BaseDmSqlParserListener) ExitEncryption_option(ctx *Encryption_optionContext) {}

// EnterCalc_option is called when production calc_option is entered.
func (s *BaseDmSqlParserListener) EnterCalc_option(ctx *Calc_optionContext) {}

// ExitCalc_option is called when production calc_option is exited.
func (s *BaseDmSqlParserListener) ExitCalc_option(ctx *Calc_optionContext) {}

// EnterFunc_action is called when production func_action is entered.
func (s *BaseDmSqlParserListener) EnterFunc_action(ctx *Func_actionContext) {}

// ExitFunc_action is called when production func_action is exited.
func (s *BaseDmSqlParserListener) ExitFunc_action(ctx *Func_actionContext) {}

// EnterFunc_call_options is called when production func_call_options is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_options(ctx *Func_call_optionsContext) {}

// ExitFunc_call_options is called when production func_call_options is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_options(ctx *Func_call_optionsContext) {}

// EnterFunc_call_option_list is called when production func_call_option_list is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_option_list(ctx *Func_call_option_listContext) {}

// ExitFunc_call_option_list is called when production func_call_option_list is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_option_list(ctx *Func_call_option_listContext) {}

// EnterFunc_call_option is called when production func_call_option is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_option(ctx *Func_call_optionContext) {}

// ExitFunc_call_option is called when production func_call_option is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_option(ctx *Func_call_optionContext) {}

// EnterInvoker_rights_clause_options is called when production invoker_rights_clause_options is entered.
func (s *BaseDmSqlParserListener) EnterInvoker_rights_clause_options(ctx *Invoker_rights_clause_optionsContext) {
}

// ExitInvoker_rights_clause_options is called when production invoker_rights_clause_options is exited.
func (s *BaseDmSqlParserListener) ExitInvoker_rights_clause_options(ctx *Invoker_rights_clause_optionsContext) {
}

// EnterInvoker_rights_clause is called when production invoker_rights_clause is entered.
func (s *BaseDmSqlParserListener) EnterInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// ExitInvoker_rights_clause is called when production invoker_rights_clause is exited.
func (s *BaseDmSqlParserListener) ExitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) {}

// EnterDeterministic_clause_options is called when production deterministic_clause_options is entered.
func (s *BaseDmSqlParserListener) EnterDeterministic_clause_options(ctx *Deterministic_clause_optionsContext) {
}

// ExitDeterministic_clause_options is called when production deterministic_clause_options is exited.
func (s *BaseDmSqlParserListener) ExitDeterministic_clause_options(ctx *Deterministic_clause_optionsContext) {
}

// EnterDeterministic_clause is called when production deterministic_clause is entered.
func (s *BaseDmSqlParserListener) EnterDeterministic_clause(ctx *Deterministic_clauseContext) {}

// ExitDeterministic_clause is called when production deterministic_clause is exited.
func (s *BaseDmSqlParserListener) ExitDeterministic_clause(ctx *Deterministic_clauseContext) {}

// EnterFunc_call_option2_options is called when production func_call_option2_options is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_option2_options(ctx *Func_call_option2_optionsContext) {
}

// ExitFunc_call_option2_options is called when production func_call_option2_options is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_option2_options(ctx *Func_call_option2_optionsContext) {
}

// EnterFunc_call_option_list2 is called when production func_call_option_list2 is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_option_list2(ctx *Func_call_option_list2Context) {}

// ExitFunc_call_option_list2 is called when production func_call_option_list2 is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_option_list2(ctx *Func_call_option_list2Context) {}

// EnterFunc_call_option2 is called when production func_call_option2 is entered.
func (s *BaseDmSqlParserListener) EnterFunc_call_option2(ctx *Func_call_option2Context) {}

// ExitFunc_call_option2 is called when production func_call_option2 is exited.
func (s *BaseDmSqlParserListener) ExitFunc_call_option2(ctx *Func_call_option2Context) {}

// EnterResult_cache_clause is called when production result_cache_clause is entered.
func (s *BaseDmSqlParserListener) EnterResult_cache_clause(ctx *Result_cache_clauseContext) {}

// ExitResult_cache_clause is called when production result_cache_clause is exited.
func (s *BaseDmSqlParserListener) ExitResult_cache_clause(ctx *Result_cache_clauseContext) {}

// EnterInner_fun_name is called when production inner_fun_name is entered.
func (s *BaseDmSqlParserListener) EnterInner_fun_name(ctx *Inner_fun_nameContext) {}

// ExitInner_fun_name is called when production inner_fun_name is exited.
func (s *BaseDmSqlParserListener) ExitInner_fun_name(ctx *Inner_fun_nameContext) {}

// EnterPlatform_type is called when production platform_type is entered.
func (s *BaseDmSqlParserListener) EnterPlatform_type(ctx *Platform_typeContext) {}

// ExitPlatform_type is called when production platform_type is exited.
func (s *BaseDmSqlParserListener) ExitPlatform_type(ctx *Platform_typeContext) {}

// EnterParam_def_list_option is called when production param_def_list_option is entered.
func (s *BaseDmSqlParserListener) EnterParam_def_list_option(ctx *Param_def_list_optionContext) {}

// ExitParam_def_list_option is called when production param_def_list_option is exited.
func (s *BaseDmSqlParserListener) ExitParam_def_list_option(ctx *Param_def_list_optionContext) {}

// EnterParam_def_list is called when production param_def_list is entered.
func (s *BaseDmSqlParserListener) EnterParam_def_list(ctx *Param_def_listContext) {}

// ExitParam_def_list is called when production param_def_list is exited.
func (s *BaseDmSqlParserListener) ExitParam_def_list(ctx *Param_def_listContext) {}

// EnterParam_def is called when production param_def is entered.
func (s *BaseDmSqlParserListener) EnterParam_def(ctx *Param_defContext) {}

// ExitParam_def is called when production param_def is exited.
func (s *BaseDmSqlParserListener) ExitParam_def(ctx *Param_defContext) {}

// EnterParam_in_out_option is called when production param_in_out_option is entered.
func (s *BaseDmSqlParserListener) EnterParam_in_out_option(ctx *Param_in_out_optionContext) {}

// ExitParam_in_out_option is called when production param_in_out_option is exited.
func (s *BaseDmSqlParserListener) ExitParam_in_out_option(ctx *Param_in_out_optionContext) {}

// EnterIs_as is called when production is_as is entered.
func (s *BaseDmSqlParserListener) EnterIs_as(ctx *Is_asContext) {}

// ExitIs_as is called when production is_as is exited.
func (s *BaseDmSqlParserListener) ExitIs_as(ctx *Is_asContext) {}

// EnterStat_on_org_stmt is called when production stat_on_org_stmt is entered.
func (s *BaseDmSqlParserListener) EnterStat_on_org_stmt(ctx *Stat_on_org_stmtContext) {}

// ExitStat_on_org_stmt is called when production stat_on_org_stmt is exited.
func (s *BaseDmSqlParserListener) ExitStat_on_org_stmt(ctx *Stat_on_org_stmtContext) {}

// EnterStat_size is called when production stat_size is entered.
func (s *BaseDmSqlParserListener) EnterStat_size(ctx *Stat_sizeContext) {}

// ExitStat_size is called when production stat_size is exited.
func (s *BaseDmSqlParserListener) ExitStat_size(ctx *Stat_sizeContext) {}

// EnterStat_para is called when production stat_para is entered.
func (s *BaseDmSqlParserListener) EnterStat_para(ctx *Stat_paraContext) {}

// ExitStat_para is called when production stat_para is exited.
func (s *BaseDmSqlParserListener) ExitStat_para(ctx *Stat_paraContext) {}

// EnterStat_summarize is called when production stat_summarize is entered.
func (s *BaseDmSqlParserListener) EnterStat_summarize(ctx *Stat_summarizeContext) {}

// ExitStat_summarize is called when production stat_summarize is exited.
func (s *BaseDmSqlParserListener) ExitStat_summarize(ctx *Stat_summarizeContext) {}

// EnterMstat_ex is called when production mstat_ex is entered.
func (s *BaseDmSqlParserListener) EnterMstat_ex(ctx *Mstat_exContext) {}

// ExitMstat_ex is called when production mstat_ex is exited.
func (s *BaseDmSqlParserListener) ExitMstat_ex(ctx *Mstat_exContext) {}

// EnterIndexid is called when production indexid is entered.
func (s *BaseDmSqlParserListener) EnterIndexid(ctx *IndexidContext) {}

// ExitIndexid is called when production indexid is exited.
func (s *BaseDmSqlParserListener) ExitIndexid(ctx *IndexidContext) {}

// EnterGlobal_tag is called when production global_tag is entered.
func (s *BaseDmSqlParserListener) EnterGlobal_tag(ctx *Global_tagContext) {}

// ExitGlobal_tag is called when production global_tag is exited.
func (s *BaseDmSqlParserListener) ExitGlobal_tag(ctx *Global_tagContext) {}

// EnterBm_join_index_clause is called when production bm_join_index_clause is entered.
func (s *BaseDmSqlParserListener) EnterBm_join_index_clause(ctx *Bm_join_index_clauseContext) {}

// ExitBm_join_index_clause is called when production bm_join_index_clause is exited.
func (s *BaseDmSqlParserListener) ExitBm_join_index_clause(ctx *Bm_join_index_clauseContext) {}

// EnterParallel_stmt is called when production parallel_stmt is entered.
func (s *BaseDmSqlParserListener) EnterParallel_stmt(ctx *Parallel_stmtContext) {}

// ExitParallel_stmt is called when production parallel_stmt is exited.
func (s *BaseDmSqlParserListener) ExitParallel_stmt(ctx *Parallel_stmtContext) {}

// EnterCreate_index_stmt is called when production create_index_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_index_stmt(ctx *Create_index_stmtContext) {}

// ExitCreate_index_stmt is called when production create_index_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_index_stmt(ctx *Create_index_stmtContext) {}

// EnterWith_inner is called when production with_inner is entered.
func (s *BaseDmSqlParserListener) EnterWith_inner(ctx *With_innerContext) {}

// ExitWith_inner is called when production with_inner is exited.
func (s *BaseDmSqlParserListener) ExitWith_inner(ctx *With_innerContext) {}

// EnterIndex_no_sort is called when production index_no_sort is entered.
func (s *BaseDmSqlParserListener) EnterIndex_no_sort(ctx *Index_no_sortContext) {}

// ExitIndex_no_sort is called when production index_no_sort is exited.
func (s *BaseDmSqlParserListener) ExitIndex_no_sort(ctx *Index_no_sortContext) {}

// EnterOnline_options is called when production online_options is entered.
func (s *BaseDmSqlParserListener) EnterOnline_options(ctx *Online_optionsContext) {}

// ExitOnline_options is called when production online_options is exited.
func (s *BaseDmSqlParserListener) ExitOnline_options(ctx *Online_optionsContext) {}

// EnterUnusable_options is called when production unusable_options is entered.
func (s *BaseDmSqlParserListener) EnterUnusable_options(ctx *Unusable_optionsContext) {}

// ExitUnusable_options is called when production unusable_options is exited.
func (s *BaseDmSqlParserListener) ExitUnusable_options(ctx *Unusable_optionsContext) {}

// EnterReverse_options is called when production reverse_options is entered.
func (s *BaseDmSqlParserListener) EnterReverse_options(ctx *Reverse_optionsContext) {}

// ExitReverse_options is called when production reverse_options is exited.
func (s *BaseDmSqlParserListener) ExitReverse_options(ctx *Reverse_optionsContext) {}

// EnterIndex_column_list is called when production index_column_list is entered.
func (s *BaseDmSqlParserListener) EnterIndex_column_list(ctx *Index_column_listContext) {}

// ExitIndex_column_list is called when production index_column_list is exited.
func (s *BaseDmSqlParserListener) ExitIndex_column_list(ctx *Index_column_listContext) {}

// EnterIndex_column_name is called when production index_column_name is entered.
func (s *BaseDmSqlParserListener) EnterIndex_column_name(ctx *Index_column_nameContext) {}

// ExitIndex_column_name is called when production index_column_name is exited.
func (s *BaseDmSqlParserListener) ExitIndex_column_name(ctx *Index_column_nameContext) {}

// EnterStorage_hash_tag is called when production storage_hash_tag is entered.
func (s *BaseDmSqlParserListener) EnterStorage_hash_tag(ctx *Storage_hash_tagContext) {}

// ExitStorage_hash_tag is called when production storage_hash_tag is exited.
func (s *BaseDmSqlParserListener) ExitStorage_hash_tag(ctx *Storage_hash_tagContext) {}

// EnterStorage_hash is called when production storage_hash is entered.
func (s *BaseDmSqlParserListener) EnterStorage_hash(ctx *Storage_hashContext) {}

// ExitStorage_hash is called when production storage_hash is exited.
func (s *BaseDmSqlParserListener) ExitStorage_hash(ctx *Storage_hashContext) {}

// EnterStorage_tag is called when production storage_tag is entered.
func (s *BaseDmSqlParserListener) EnterStorage_tag(ctx *Storage_tagContext) {}

// ExitStorage_tag is called when production storage_tag is exited.
func (s *BaseDmSqlParserListener) ExitStorage_tag(ctx *Storage_tagContext) {}

// EnterStorage_tag_nn is called when production storage_tag_nn is entered.
func (s *BaseDmSqlParserListener) EnterStorage_tag_nn(ctx *Storage_tag_nnContext) {}

// ExitStorage_tag_nn is called when production storage_tag_nn is exited.
func (s *BaseDmSqlParserListener) ExitStorage_tag_nn(ctx *Storage_tag_nnContext) {}

// EnterTablespace_clause is called when production tablespace_clause is entered.
func (s *BaseDmSqlParserListener) EnterTablespace_clause(ctx *Tablespace_clauseContext) {}

// ExitTablespace_clause is called when production tablespace_clause is exited.
func (s *BaseDmSqlParserListener) ExitTablespace_clause(ctx *Tablespace_clauseContext) {}

// EnterObject_table_substitution_clause is called when production object_table_substitution_clause is entered.
func (s *BaseDmSqlParserListener) EnterObject_table_substitution_clause(ctx *Object_table_substitution_clauseContext) {
}

// ExitObject_table_substitution_clause is called when production object_table_substitution_clause is exited.
func (s *BaseDmSqlParserListener) ExitObject_table_substitution_clause(ctx *Object_table_substitution_clauseContext) {
}

// EnterObject_table_substitution is called when production object_table_substitution is entered.
func (s *BaseDmSqlParserListener) EnterObject_table_substitution(ctx *Object_table_substitutionContext) {
}

// ExitObject_table_substitution is called when production object_table_substitution is exited.
func (s *BaseDmSqlParserListener) ExitObject_table_substitution(ctx *Object_table_substitutionContext) {
}

// EnterOid_clause is called when production oid_clause is entered.
func (s *BaseDmSqlParserListener) EnterOid_clause(ctx *Oid_clauseContext) {}

// ExitOid_clause is called when production oid_clause is exited.
func (s *BaseDmSqlParserListener) ExitOid_clause(ctx *Oid_clauseContext) {}

// EnterOid_gen_type is called when production oid_gen_type is entered.
func (s *BaseDmSqlParserListener) EnterOid_gen_type(ctx *Oid_gen_typeContext) {}

// ExitOid_gen_type is called when production oid_gen_type is exited.
func (s *BaseDmSqlParserListener) ExitOid_gen_type(ctx *Oid_gen_typeContext) {}

// EnterOid_index_clause is called when production oid_index_clause is entered.
func (s *BaseDmSqlParserListener) EnterOid_index_clause(ctx *Oid_index_clauseContext) {}

// ExitOid_index_clause is called when production oid_index_clause is exited.
func (s *BaseDmSqlParserListener) ExitOid_index_clause(ctx *Oid_index_clauseContext) {}

// EnterOid_tablespace_clause is called when production oid_tablespace_clause is entered.
func (s *BaseDmSqlParserListener) EnterOid_tablespace_clause(ctx *Oid_tablespace_clauseContext) {}

// ExitOid_tablespace_clause is called when production oid_tablespace_clause is exited.
func (s *BaseDmSqlParserListener) ExitOid_tablespace_clause(ctx *Oid_tablespace_clauseContext) {}

// EnterOid_tablespace_name is called when production oid_tablespace_name is entered.
func (s *BaseDmSqlParserListener) EnterOid_tablespace_name(ctx *Oid_tablespace_nameContext) {}

// ExitOid_tablespace_name is called when production oid_tablespace_name is exited.
func (s *BaseDmSqlParserListener) ExitOid_tablespace_name(ctx *Oid_tablespace_nameContext) {}

// EnterLocal_option is called when production local_option is entered.
func (s *BaseDmSqlParserListener) EnterLocal_option(ctx *Local_optionContext) {}

// ExitLocal_option is called when production local_option is exited.
func (s *BaseDmSqlParserListener) ExitLocal_option(ctx *Local_optionContext) {}

// EnterStorage_list is called when production storage_list is entered.
func (s *BaseDmSqlParserListener) EnterStorage_list(ctx *Storage_listContext) {}

// ExitStorage_list is called when production storage_list is exited.
func (s *BaseDmSqlParserListener) ExitStorage_list(ctx *Storage_listContext) {}

// EnterStorage_hashpartmap is called when production storage_hashpartmap is entered.
func (s *BaseDmSqlParserListener) EnterStorage_hashpartmap(ctx *Storage_hashpartmapContext) {}

// ExitStorage_hashpartmap is called when production storage_hashpartmap is exited.
func (s *BaseDmSqlParserListener) ExitStorage_hashpartmap(ctx *Storage_hashpartmapContext) {}

// EnterStorage is called when production storage is entered.
func (s *BaseDmSqlParserListener) EnterStorage(ctx *StorageContext) {}

// ExitStorage is called when production storage is exited.
func (s *BaseDmSqlParserListener) ExitStorage(ctx *StorageContext) {}

// EnterId_list is called when production id_list is entered.
func (s *BaseDmSqlParserListener) EnterId_list(ctx *Id_listContext) {}

// ExitId_list is called when production id_list is exited.
func (s *BaseDmSqlParserListener) ExitId_list(ctx *Id_listContext) {}

// EnterCreate_proc_stmt is called when production create_proc_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_proc_stmt(ctx *Create_proc_stmtContext) {}

// ExitCreate_proc_stmt is called when production create_proc_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_proc_stmt(ctx *Create_proc_stmtContext) {}

// EnterCreate_package_stmt is called when production create_package_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_package_stmt(ctx *Create_package_stmtContext) {}

// ExitCreate_package_stmt is called when production create_package_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_package_stmt(ctx *Create_package_stmtContext) {}

// EnterPkg_cls_flag is called when production pkg_cls_flag is entered.
func (s *BaseDmSqlParserListener) EnterPkg_cls_flag(ctx *Pkg_cls_flagContext) {}

// ExitPkg_cls_flag is called when production pkg_cls_flag is exited.
func (s *BaseDmSqlParserListener) ExitPkg_cls_flag(ctx *Pkg_cls_flagContext) {}

// EnterBlk_end_option is called when production blk_end_option is entered.
func (s *BaseDmSqlParserListener) EnterBlk_end_option(ctx *Blk_end_optionContext) {}

// ExitBlk_end_option is called when production blk_end_option is exited.
func (s *BaseDmSqlParserListener) ExitBlk_end_option(ctx *Blk_end_optionContext) {}

// EnterBlk_end_option_low is called when production blk_end_option_low is entered.
func (s *BaseDmSqlParserListener) EnterBlk_end_option_low(ctx *Blk_end_option_lowContext) {}

// ExitBlk_end_option_low is called when production blk_end_option_low is exited.
func (s *BaseDmSqlParserListener) ExitBlk_end_option_low(ctx *Blk_end_option_lowContext) {}

// EnterPackage_def_list_options is called when production package_def_list_options is entered.
func (s *BaseDmSqlParserListener) EnterPackage_def_list_options(ctx *Package_def_list_optionsContext) {
}

// ExitPackage_def_list_options is called when production package_def_list_options is exited.
func (s *BaseDmSqlParserListener) ExitPackage_def_list_options(ctx *Package_def_list_optionsContext) {
}

// EnterPackage_def_list is called when production package_def_list is entered.
func (s *BaseDmSqlParserListener) EnterPackage_def_list(ctx *Package_def_listContext) {}

// ExitPackage_def_list is called when production package_def_list is exited.
func (s *BaseDmSqlParserListener) ExitPackage_def_list(ctx *Package_def_listContext) {}

// EnterPackage_def is called when production package_def is entered.
func (s *BaseDmSqlParserListener) EnterPackage_def(ctx *Package_defContext) {}

// ExitPackage_def is called when production package_def is exited.
func (s *BaseDmSqlParserListener) ExitPackage_def(ctx *Package_defContext) {}

// EnterRestrict_param_lst is called when production restrict_param_lst is entered.
func (s *BaseDmSqlParserListener) EnterRestrict_param_lst(ctx *Restrict_param_lstContext) {}

// ExitRestrict_param_lst is called when production restrict_param_lst is exited.
func (s *BaseDmSqlParserListener) ExitRestrict_param_lst(ctx *Restrict_param_lstContext) {}

// EnterCreate_package_body_stmt is called when production create_package_body_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_package_body_stmt(ctx *Create_package_body_stmtContext) {
}

// ExitCreate_package_body_stmt is called when production create_package_body_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_package_body_stmt(ctx *Create_package_body_stmtContext) {
}

// EnterCreate_pkg_body_header is called when production create_pkg_body_header is entered.
func (s *BaseDmSqlParserListener) EnterCreate_pkg_body_header(ctx *Create_pkg_body_headerContext) {}

// ExitCreate_pkg_body_header is called when production create_pkg_body_header is exited.
func (s *BaseDmSqlParserListener) ExitCreate_pkg_body_header(ctx *Create_pkg_body_headerContext) {}

// EnterPkg_cls_body_flag is called when production pkg_cls_body_flag is entered.
func (s *BaseDmSqlParserListener) EnterPkg_cls_body_flag(ctx *Pkg_cls_body_flagContext) {}

// ExitPkg_cls_body_flag is called when production pkg_cls_body_flag is exited.
func (s *BaseDmSqlParserListener) ExitPkg_cls_body_flag(ctx *Pkg_cls_body_flagContext) {}

// EnterPackage_body_init_option is called when production package_body_init_option is entered.
func (s *BaseDmSqlParserListener) EnterPackage_body_init_option(ctx *Package_body_init_optionContext) {
}

// ExitPackage_body_init_option is called when production package_body_init_option is exited.
func (s *BaseDmSqlParserListener) ExitPackage_body_init_option(ctx *Package_body_init_optionContext) {
}

// EnterPackage_body_def_list is called when production package_body_def_list is entered.
func (s *BaseDmSqlParserListener) EnterPackage_body_def_list(ctx *Package_body_def_listContext) {}

// ExitPackage_body_def_list is called when production package_body_def_list is exited.
func (s *BaseDmSqlParserListener) ExitPackage_body_def_list(ctx *Package_body_def_listContext) {}

// EnterPackage_body_def is called when production package_body_def is entered.
func (s *BaseDmSqlParserListener) EnterPackage_body_def(ctx *Package_body_defContext) {}

// ExitPackage_body_def is called when production package_body_def is exited.
func (s *BaseDmSqlParserListener) ExitPackage_body_def(ctx *Package_body_defContext) {}

// EnterPackage_body_def2 is called when production package_body_def2 is entered.
func (s *BaseDmSqlParserListener) EnterPackage_body_def2(ctx *Package_body_def2Context) {}

// ExitPackage_body_def2 is called when production package_body_def2 is exited.
func (s *BaseDmSqlParserListener) ExitPackage_body_def2(ctx *Package_body_def2Context) {}

// EnterCheck_exec_options is called when production check_exec_options is entered.
func (s *BaseDmSqlParserListener) EnterCheck_exec_options(ctx *Check_exec_optionsContext) {}

// ExitCheck_exec_options is called when production check_exec_options is exited.
func (s *BaseDmSqlParserListener) ExitCheck_exec_options(ctx *Check_exec_optionsContext) {}

// EnterSubpg_decl_stmt is called when production subpg_decl_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSubpg_decl_stmt(ctx *Subpg_decl_stmtContext) {}

// ExitSubpg_decl_stmt is called when production subpg_decl_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSubpg_decl_stmt(ctx *Subpg_decl_stmtContext) {}

// EnterCreate_type_stmt is called when production create_type_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_type_stmt(ctx *Create_type_stmtContext) {}

// ExitCreate_type_stmt is called when production create_type_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_type_stmt(ctx *Create_type_stmtContext) {}

// EnterForce_option is called when production force_option is entered.
func (s *BaseDmSqlParserListener) EnterForce_option(ctx *Force_optionContext) {}

// ExitForce_option is called when production force_option is exited.
func (s *BaseDmSqlParserListener) ExitForce_option(ctx *Force_optionContext) {}

// EnterObject_option is called when production object_option is entered.
func (s *BaseDmSqlParserListener) EnterObject_option(ctx *Object_optionContext) {}

// ExitObject_option is called when production object_option is exited.
func (s *BaseDmSqlParserListener) ExitObject_option(ctx *Object_optionContext) {}

// EnterUnder_option is called when production under_option is entered.
func (s *BaseDmSqlParserListener) EnterUnder_option(ctx *Under_optionContext) {}

// ExitUnder_option is called when production under_option is exited.
func (s *BaseDmSqlParserListener) ExitUnder_option(ctx *Under_optionContext) {}

// EnterObject_def_list_options is called when production object_def_list_options is entered.
func (s *BaseDmSqlParserListener) EnterObject_def_list_options(ctx *Object_def_list_optionsContext) {}

// ExitObject_def_list_options is called when production object_def_list_options is exited.
func (s *BaseDmSqlParserListener) ExitObject_def_list_options(ctx *Object_def_list_optionsContext) {}

// EnterObject_def_list is called when production object_def_list is entered.
func (s *BaseDmSqlParserListener) EnterObject_def_list(ctx *Object_def_listContext) {}

// ExitObject_def_list is called when production object_def_list is exited.
func (s *BaseDmSqlParserListener) ExitObject_def_list(ctx *Object_def_listContext) {}

// EnterObject_def is called when production object_def is entered.
func (s *BaseDmSqlParserListener) EnterObject_def(ctx *Object_defContext) {}

// ExitObject_def is called when production object_def is exited.
func (s *BaseDmSqlParserListener) ExitObject_def(ctx *Object_defContext) {}

// EnterMember_static is called when production member_static is entered.
func (s *BaseDmSqlParserListener) EnterMember_static(ctx *Member_staticContext) {}

// ExitMember_static is called when production member_static is exited.
func (s *BaseDmSqlParserListener) ExitMember_static(ctx *Member_staticContext) {}

// EnterMethod_inherit_options is called when production method_inherit_options is entered.
func (s *BaseDmSqlParserListener) EnterMethod_inherit_options(ctx *Method_inherit_optionsContext) {}

// ExitMethod_inherit_options is called when production method_inherit_options is exited.
func (s *BaseDmSqlParserListener) ExitMethod_inherit_options(ctx *Method_inherit_optionsContext) {}

// EnterMethod_inherit_option is called when production method_inherit_option is entered.
func (s *BaseDmSqlParserListener) EnterMethod_inherit_option(ctx *Method_inherit_optionContext) {}

// ExitMethod_inherit_option is called when production method_inherit_option is exited.
func (s *BaseDmSqlParserListener) ExitMethod_inherit_option(ctx *Method_inherit_optionContext) {}

// EnterFinal_inst_list_options is called when production final_inst_list_options is entered.
func (s *BaseDmSqlParserListener) EnterFinal_inst_list_options(ctx *Final_inst_list_optionsContext) {}

// ExitFinal_inst_list_options is called when production final_inst_list_options is exited.
func (s *BaseDmSqlParserListener) ExitFinal_inst_list_options(ctx *Final_inst_list_optionsContext) {}

// EnterFinal_inst_list is called when production final_inst_list is entered.
func (s *BaseDmSqlParserListener) EnterFinal_inst_list(ctx *Final_inst_listContext) {}

// ExitFinal_inst_list is called when production final_inst_list is exited.
func (s *BaseDmSqlParserListener) ExitFinal_inst_list(ctx *Final_inst_listContext) {}

// EnterFinal_inst is called when production final_inst is entered.
func (s *BaseDmSqlParserListener) EnterFinal_inst(ctx *Final_instContext) {}

// ExitFinal_inst is called when production final_inst is exited.
func (s *BaseDmSqlParserListener) ExitFinal_inst(ctx *Final_instContext) {}

// EnterOverriding_option is called when production overriding_option is entered.
func (s *BaseDmSqlParserListener) EnterOverriding_option(ctx *Overriding_optionContext) {}

// ExitOverriding_option is called when production overriding_option is exited.
func (s *BaseDmSqlParserListener) ExitOverriding_option(ctx *Overriding_optionContext) {}

// EnterMethod_attr_options is called when production method_attr_options is entered.
func (s *BaseDmSqlParserListener) EnterMethod_attr_options(ctx *Method_attr_optionsContext) {}

// ExitMethod_attr_options is called when production method_attr_options is exited.
func (s *BaseDmSqlParserListener) ExitMethod_attr_options(ctx *Method_attr_optionsContext) {}

// EnterMethod_attr is called when production method_attr is entered.
func (s *BaseDmSqlParserListener) EnterMethod_attr(ctx *Method_attrContext) {}

// ExitMethod_attr is called when production method_attr is exited.
func (s *BaseDmSqlParserListener) ExitMethod_attr(ctx *Method_attrContext) {}

// EnterCreate_type_body_stmt is called when production create_type_body_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_type_body_stmt(ctx *Create_type_body_stmtContext) {}

// ExitCreate_type_body_stmt is called when production create_type_body_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_type_body_stmt(ctx *Create_type_body_stmtContext) {}

// EnterObject_body_def_list is called when production object_body_def_list is entered.
func (s *BaseDmSqlParserListener) EnterObject_body_def_list(ctx *Object_body_def_listContext) {}

// ExitObject_body_def_list is called when production object_body_def_list is exited.
func (s *BaseDmSqlParserListener) ExitObject_body_def_list(ctx *Object_body_def_listContext) {}

// EnterObject_body_def is called when production object_body_def is entered.
func (s *BaseDmSqlParserListener) EnterObject_body_def(ctx *Object_body_defContext) {}

// ExitObject_body_def is called when production object_body_def is exited.
func (s *BaseDmSqlParserListener) ExitObject_body_def(ctx *Object_body_defContext) {}

// EnterCreate_context_stmt is called when production create_context_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_context_stmt(ctx *Create_context_stmtContext) {}

// ExitCreate_context_stmt is called when production create_context_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_context_stmt(ctx *Create_context_stmtContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseDmSqlParserListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseDmSqlParserListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterInitialized is called when production initialized is entered.
func (s *BaseDmSqlParserListener) EnterInitialized(ctx *InitializedContext) {}

// ExitInitialized is called when production initialized is exited.
func (s *BaseDmSqlParserListener) ExitInitialized(ctx *InitializedContext) {}

// EnterCreate_directory_stmt is called when production create_directory_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_directory_stmt(ctx *Create_directory_stmtContext) {}

// ExitCreate_directory_stmt is called when production create_directory_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_directory_stmt(ctx *Create_directory_stmtContext) {}

// EnterCreate_crypto_stmt is called when production create_crypto_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_crypto_stmt(ctx *Create_crypto_stmtContext) {}

// ExitCreate_crypto_stmt is called when production create_crypto_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_crypto_stmt(ctx *Create_crypto_stmtContext) {}

// EnterAlter_crypto_stmt is called when production alter_crypto_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_crypto_stmt(ctx *Alter_crypto_stmtContext) {}

// ExitAlter_crypto_stmt is called when production alter_crypto_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_crypto_stmt(ctx *Alter_crypto_stmtContext) {}

// EnterAlter_crypto_action is called when production alter_crypto_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_crypto_action(ctx *Alter_crypto_actionContext) {}

// ExitAlter_crypto_action is called when production alter_crypto_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_crypto_action(ctx *Alter_crypto_actionContext) {}

// EnterComment_stmt is called when production comment_stmt is entered.
func (s *BaseDmSqlParserListener) EnterComment_stmt(ctx *Comment_stmtContext) {}

// ExitComment_stmt is called when production comment_stmt is exited.
func (s *BaseDmSqlParserListener) ExitComment_stmt(ctx *Comment_stmtContext) {}

// EnterComment_tag is called when production comment_tag is entered.
func (s *BaseDmSqlParserListener) EnterComment_tag(ctx *Comment_tagContext) {}

// ExitComment_tag is called when production comment_tag is exited.
func (s *BaseDmSqlParserListener) ExitComment_tag(ctx *Comment_tagContext) {}

// EnterCreate_partition_group_stmt is called when production create_partition_group_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_partition_group_stmt(ctx *Create_partition_group_stmtContext) {
}

// ExitCreate_partition_group_stmt is called when production create_partition_group_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_partition_group_stmt(ctx *Create_partition_group_stmtContext) {
}

// EnterStorage_act_datatype is called when production storage_act_datatype is entered.
func (s *BaseDmSqlParserListener) EnterStorage_act_datatype(ctx *Storage_act_datatypeContext) {}

// ExitStorage_act_datatype is called when production storage_act_datatype is exited.
func (s *BaseDmSqlParserListener) ExitStorage_act_datatype(ctx *Storage_act_datatypeContext) {}

// EnterPg_storage_lst is called when production pg_storage_lst is entered.
func (s *BaseDmSqlParserListener) EnterPg_storage_lst(ctx *Pg_storage_lstContext) {}

// ExitPg_storage_lst is called when production pg_storage_lst is exited.
func (s *BaseDmSqlParserListener) ExitPg_storage_lst(ctx *Pg_storage_lstContext) {}

// EnterPg_storage is called when production pg_storage is entered.
func (s *BaseDmSqlParserListener) EnterPg_storage(ctx *Pg_storageContext) {}

// ExitPg_storage is called when production pg_storage is exited.
func (s *BaseDmSqlParserListener) ExitPg_storage(ctx *Pg_storageContext) {}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterCtab_append_attr_clause is called when production ctab_append_attr_clause is entered.
func (s *BaseDmSqlParserListener) EnterCtab_append_attr_clause(ctx *Ctab_append_attr_clauseContext) {}

// ExitCtab_append_attr_clause is called when production ctab_append_attr_clause is exited.
func (s *BaseDmSqlParserListener) ExitCtab_append_attr_clause(ctx *Ctab_append_attr_clauseContext) {}

// EnterCtab_append_attr_list is called when production ctab_append_attr_list is entered.
func (s *BaseDmSqlParserListener) EnterCtab_append_attr_list(ctx *Ctab_append_attr_listContext) {}

// ExitCtab_append_attr_list is called when production ctab_append_attr_list is exited.
func (s *BaseDmSqlParserListener) ExitCtab_append_attr_list(ctx *Ctab_append_attr_listContext) {}

// EnterCobjtab_append_attr_clause is called when production cobjtab_append_attr_clause is entered.
func (s *BaseDmSqlParserListener) EnterCobjtab_append_attr_clause(ctx *Cobjtab_append_attr_clauseContext) {
}

// ExitCobjtab_append_attr_clause is called when production cobjtab_append_attr_clause is exited.
func (s *BaseDmSqlParserListener) ExitCobjtab_append_attr_clause(ctx *Cobjtab_append_attr_clauseContext) {
}

// EnterCobjtab_append_attr_list is called when production cobjtab_append_attr_list is entered.
func (s *BaseDmSqlParserListener) EnterCobjtab_append_attr_list(ctx *Cobjtab_append_attr_listContext) {
}

// ExitCobjtab_append_attr_list is called when production cobjtab_append_attr_list is exited.
func (s *BaseDmSqlParserListener) ExitCobjtab_append_attr_list(ctx *Cobjtab_append_attr_listContext) {
}

// EnterCtab_append_attr is called when production ctab_append_attr is entered.
func (s *BaseDmSqlParserListener) EnterCtab_append_attr(ctx *Ctab_append_attrContext) {}

// ExitCtab_append_attr is called when production ctab_append_attr is exited.
func (s *BaseDmSqlParserListener) ExitCtab_append_attr(ctx *Ctab_append_attrContext) {}

// EnterCobjtab_append_attr is called when production cobjtab_append_attr is entered.
func (s *BaseDmSqlParserListener) EnterCobjtab_append_attr(ctx *Cobjtab_append_attrContext) {}

// ExitCobjtab_append_attr is called when production cobjtab_append_attr is exited.
func (s *BaseDmSqlParserListener) ExitCobjtab_append_attr(ctx *Cobjtab_append_attrContext) {}

// EnterCreate_table_action is called when production create_table_action is entered.
func (s *BaseDmSqlParserListener) EnterCreate_table_action(ctx *Create_table_actionContext) {}

// ExitCreate_table_action is called when production create_table_action is exited.
func (s *BaseDmSqlParserListener) ExitCreate_table_action(ctx *Create_table_actionContext) {}

// EnterCtab_log_options is called when production ctab_log_options is entered.
func (s *BaseDmSqlParserListener) EnterCtab_log_options(ctx *Ctab_log_optionsContext) {}

// ExitCtab_log_options is called when production ctab_log_options is exited.
func (s *BaseDmSqlParserListener) ExitCtab_log_options(ctx *Ctab_log_optionsContext) {}

// EnterCtab_log_option is called when production ctab_log_option is entered.
func (s *BaseDmSqlParserListener) EnterCtab_log_option(ctx *Ctab_log_optionContext) {}

// ExitCtab_log_option is called when production ctab_log_option is exited.
func (s *BaseDmSqlParserListener) ExitCtab_log_option(ctx *Ctab_log_optionContext) {}

// EnterCtab_error_options is called when production ctab_error_options is entered.
func (s *BaseDmSqlParserListener) EnterCtab_error_options(ctx *Ctab_error_optionsContext) {}

// ExitCtab_error_options is called when production ctab_error_options is exited.
func (s *BaseDmSqlParserListener) ExitCtab_error_options(ctx *Ctab_error_optionsContext) {}

// EnterAdvance_log_clause is called when production advance_log_clause is entered.
func (s *BaseDmSqlParserListener) EnterAdvance_log_clause(ctx *Advance_log_clauseContext) {}

// ExitAdvance_log_clause is called when production advance_log_clause is exited.
func (s *BaseDmSqlParserListener) ExitAdvance_log_clause(ctx *Advance_log_clauseContext) {}

// EnterAdd_log_clause is called when production add_log_clause is entered.
func (s *BaseDmSqlParserListener) EnterAdd_log_clause(ctx *Add_log_clauseContext) {}

// ExitAdd_log_clause is called when production add_log_clause is exited.
func (s *BaseDmSqlParserListener) ExitAdd_log_clause(ctx *Add_log_clauseContext) {}

// EnterCtab_error_option is called when production ctab_error_option is entered.
func (s *BaseDmSqlParserListener) EnterCtab_error_option(ctx *Ctab_error_optionContext) {}

// ExitCtab_error_option is called when production ctab_error_option is exited.
func (s *BaseDmSqlParserListener) ExitCtab_error_option(ctx *Ctab_error_optionContext) {}

// EnterCtab_row_movement_clause is called when production ctab_row_movement_clause is entered.
func (s *BaseDmSqlParserListener) EnterCtab_row_movement_clause(ctx *Ctab_row_movement_clauseContext) {
}

// ExitCtab_row_movement_clause is called when production ctab_row_movement_clause is exited.
func (s *BaseDmSqlParserListener) ExitCtab_row_movement_clause(ctx *Ctab_row_movement_clauseContext) {
}

// EnterRange_distribute_act is called when production range_distribute_act is entered.
func (s *BaseDmSqlParserListener) EnterRange_distribute_act(ctx *Range_distribute_actContext) {}

// ExitRange_distribute_act is called when production range_distribute_act is exited.
func (s *BaseDmSqlParserListener) ExitRange_distribute_act(ctx *Range_distribute_actContext) {}

// EnterRange_distribute_act_lst is called when production range_distribute_act_lst is entered.
func (s *BaseDmSqlParserListener) EnterRange_distribute_act_lst(ctx *Range_distribute_act_lstContext) {
}

// ExitRange_distribute_act_lst is called when production range_distribute_act_lst is exited.
func (s *BaseDmSqlParserListener) ExitRange_distribute_act_lst(ctx *Range_distribute_act_lstContext) {
}

// EnterList_distribute_act is called when production list_distribute_act is entered.
func (s *BaseDmSqlParserListener) EnterList_distribute_act(ctx *List_distribute_actContext) {}

// ExitList_distribute_act is called when production list_distribute_act is exited.
func (s *BaseDmSqlParserListener) ExitList_distribute_act(ctx *List_distribute_actContext) {}

// EnterList_distribute_act_list is called when production list_distribute_act_list is entered.
func (s *BaseDmSqlParserListener) EnterList_distribute_act_list(ctx *List_distribute_act_listContext) {
}

// ExitList_distribute_act_list is called when production list_distribute_act_list is exited.
func (s *BaseDmSqlParserListener) ExitList_distribute_act_list(ctx *List_distribute_act_listContext) {
}

// EnterDistribute_by_option is called when production distribute_by_option is entered.
func (s *BaseDmSqlParserListener) EnterDistribute_by_option(ctx *Distribute_by_optionContext) {}

// ExitDistribute_by_option is called when production distribute_by_option is exited.
func (s *BaseDmSqlParserListener) ExitDistribute_by_option(ctx *Distribute_by_optionContext) {}

// EnterDistribute_by is called when production distribute_by is entered.
func (s *BaseDmSqlParserListener) EnterDistribute_by(ctx *Distribute_byContext) {}

// ExitDistribute_by is called when production distribute_by is exited.
func (s *BaseDmSqlParserListener) ExitDistribute_by(ctx *Distribute_byContext) {}

// EnterIncrement_set is called when production increment_set is entered.
func (s *BaseDmSqlParserListener) EnterIncrement_set(ctx *Increment_setContext) {}

// ExitIncrement_set is called when production increment_set is exited.
func (s *BaseDmSqlParserListener) ExitIncrement_set(ctx *Increment_setContext) {}

// EnterIncrement is called when production increment is entered.
func (s *BaseDmSqlParserListener) EnterIncrement(ctx *IncrementContext) {}

// ExitIncrement is called when production increment is exited.
func (s *BaseDmSqlParserListener) ExitIncrement(ctx *IncrementContext) {}

// EnterRowdependencies_clause is called when production rowdependencies_clause is entered.
func (s *BaseDmSqlParserListener) EnterRowdependencies_clause(ctx *Rowdependencies_clauseContext) {}

// ExitRowdependencies_clause is called when production rowdependencies_clause is exited.
func (s *BaseDmSqlParserListener) ExitRowdependencies_clause(ctx *Rowdependencies_clauseContext) {}

// EnterPg_sub_partition is called when production pg_sub_partition is entered.
func (s *BaseDmSqlParserListener) EnterPg_sub_partition(ctx *Pg_sub_partitionContext) {}

// ExitPg_sub_partition is called when production pg_sub_partition is exited.
func (s *BaseDmSqlParserListener) ExitPg_sub_partition(ctx *Pg_sub_partitionContext) {}

// EnterTable_type_option is called when production table_type_option is entered.
func (s *BaseDmSqlParserListener) EnterTable_type_option(ctx *Table_type_optionContext) {}

// ExitTable_type_option is called when production table_type_option is exited.
func (s *BaseDmSqlParserListener) ExitTable_type_option(ctx *Table_type_optionContext) {}

// EnterTable_temp_option is called when production table_temp_option is entered.
func (s *BaseDmSqlParserListener) EnterTable_temp_option(ctx *Table_temp_optionContext) {}

// ExitTable_temp_option is called when production table_temp_option is exited.
func (s *BaseDmSqlParserListener) ExitTable_temp_option(ctx *Table_temp_optionContext) {}

// EnterObjtab_elem_constraint is called when production objtab_elem_constraint is entered.
func (s *BaseDmSqlParserListener) EnterObjtab_elem_constraint(ctx *Objtab_elem_constraintContext) {}

// ExitObjtab_elem_constraint is called when production objtab_elem_constraint is exited.
func (s *BaseDmSqlParserListener) ExitObjtab_elem_constraint(ctx *Objtab_elem_constraintContext) {}

// EnterObjtab_element_cons_list is called when production objtab_element_cons_list is entered.
func (s *BaseDmSqlParserListener) EnterObjtab_element_cons_list(ctx *Objtab_element_cons_listContext) {
}

// ExitObjtab_element_cons_list is called when production objtab_element_cons_list is exited.
func (s *BaseDmSqlParserListener) ExitObjtab_element_cons_list(ctx *Objtab_element_cons_listContext) {
}

// EnterObjtab_element_cons is called when production objtab_element_cons is entered.
func (s *BaseDmSqlParserListener) EnterObjtab_element_cons(ctx *Objtab_element_consContext) {}

// ExitObjtab_element_cons is called when production objtab_element_cons is exited.
func (s *BaseDmSqlParserListener) ExitObjtab_element_cons(ctx *Objtab_element_consContext) {}

// EnterObjcol_constraint is called when production objcol_constraint is entered.
func (s *BaseDmSqlParserListener) EnterObjcol_constraint(ctx *Objcol_constraintContext) {}

// ExitObjcol_constraint is called when production objcol_constraint is exited.
func (s *BaseDmSqlParserListener) ExitObjcol_constraint(ctx *Objcol_constraintContext) {}

// EnterTable_element_list is called when production table_element_list is entered.
func (s *BaseDmSqlParserListener) EnterTable_element_list(ctx *Table_element_listContext) {}

// ExitTable_element_list is called when production table_element_list is exited.
func (s *BaseDmSqlParserListener) ExitTable_element_list(ctx *Table_element_listContext) {}

// EnterTable_element is called when production table_element is entered.
func (s *BaseDmSqlParserListener) EnterTable_element(ctx *Table_elementContext) {}

// ExitTable_element is called when production table_element is exited.
func (s *BaseDmSqlParserListener) ExitTable_element(ctx *Table_elementContext) {}

// EnterTable_constraint_def is called when production table_constraint_def is entered.
func (s *BaseDmSqlParserListener) EnterTable_constraint_def(ctx *Table_constraint_defContext) {}

// ExitTable_constraint_def is called when production table_constraint_def is exited.
func (s *BaseDmSqlParserListener) ExitTable_constraint_def(ctx *Table_constraint_defContext) {}

// EnterOn_commit_option is called when production on_commit_option is entered.
func (s *BaseDmSqlParserListener) EnterOn_commit_option(ctx *On_commit_optionContext) {}

// ExitOn_commit_option is called when production on_commit_option is exited.
func (s *BaseDmSqlParserListener) ExitOn_commit_option(ctx *On_commit_optionContext) {}

// EnterOn_commit_option_nn is called when production on_commit_option_nn is entered.
func (s *BaseDmSqlParserListener) EnterOn_commit_option_nn(ctx *On_commit_option_nnContext) {}

// ExitOn_commit_option_nn is called when production on_commit_option_nn is exited.
func (s *BaseDmSqlParserListener) ExitOn_commit_option_nn(ctx *On_commit_option_nnContext) {}

// EnterLogging_option is called when production logging_option is entered.
func (s *BaseDmSqlParserListener) EnterLogging_option(ctx *Logging_optionContext) {}

// ExitLogging_option is called when production logging_option is exited.
func (s *BaseDmSqlParserListener) ExitLogging_option(ctx *Logging_optionContext) {}

// EnterLogging_option_nn is called when production logging_option_nn is entered.
func (s *BaseDmSqlParserListener) EnterLogging_option_nn(ctx *Logging_option_nnContext) {}

// ExitLogging_option_nn is called when production logging_option_nn is exited.
func (s *BaseDmSqlParserListener) ExitLogging_option_nn(ctx *Logging_option_nnContext) {}

// EnterPartition_clause is called when production partition_clause is entered.
func (s *BaseDmSqlParserListener) EnterPartition_clause(ctx *Partition_clauseContext) {}

// ExitPartition_clause is called when production partition_clause is exited.
func (s *BaseDmSqlParserListener) ExitPartition_clause(ctx *Partition_clauseContext) {}

// EnterPartition_clause_nn is called when production partition_clause_nn is entered.
func (s *BaseDmSqlParserListener) EnterPartition_clause_nn(ctx *Partition_clause_nnContext) {}

// ExitPartition_clause_nn is called when production partition_clause_nn is exited.
func (s *BaseDmSqlParserListener) ExitPartition_clause_nn(ctx *Partition_clause_nnContext) {}

// EnterHorizon_partition_clause is called when production horizon_partition_clause is entered.
func (s *BaseDmSqlParserListener) EnterHorizon_partition_clause(ctx *Horizon_partition_clauseContext) {
}

// ExitHorizon_partition_clause is called when production horizon_partition_clause is exited.
func (s *BaseDmSqlParserListener) ExitHorizon_partition_clause(ctx *Horizon_partition_clauseContext) {
}

// EnterCompress_tag_hdr is called when production compress_tag_hdr is entered.
func (s *BaseDmSqlParserListener) EnterCompress_tag_hdr(ctx *Compress_tag_hdrContext) {}

// ExitCompress_tag_hdr is called when production compress_tag_hdr is exited.
func (s *BaseDmSqlParserListener) ExitCompress_tag_hdr(ctx *Compress_tag_hdrContext) {}

// EnterCompress_clause_opt is called when production compress_clause_opt is entered.
func (s *BaseDmSqlParserListener) EnterCompress_clause_opt(ctx *Compress_clause_optContext) {}

// ExitCompress_clause_opt is called when production compress_clause_opt is exited.
func (s *BaseDmSqlParserListener) ExitCompress_clause_opt(ctx *Compress_clause_optContext) {}

// EnterCompress_tag is called when production compress_tag is entered.
func (s *BaseDmSqlParserListener) EnterCompress_tag(ctx *Compress_tagContext) {}

// ExitCompress_tag is called when production compress_tag is exited.
func (s *BaseDmSqlParserListener) ExitCompress_tag(ctx *Compress_tagContext) {}

// EnterCompress_level is called when production compress_level is entered.
func (s *BaseDmSqlParserListener) EnterCompress_level(ctx *Compress_levelContext) {}

// ExitCompress_level is called when production compress_level is exited.
func (s *BaseDmSqlParserListener) ExitCompress_level(ctx *Compress_levelContext) {}

// EnterCompress_type is called when production compress_type is entered.
func (s *BaseDmSqlParserListener) EnterCompress_type(ctx *Compress_typeContext) {}

// ExitCompress_type is called when production compress_type is exited.
func (s *BaseDmSqlParserListener) ExitCompress_type(ctx *Compress_typeContext) {}

// EnterRange_partition is called when production range_partition is entered.
func (s *BaseDmSqlParserListener) EnterRange_partition(ctx *Range_partitionContext) {}

// ExitRange_partition is called when production range_partition is exited.
func (s *BaseDmSqlParserListener) ExitRange_partition(ctx *Range_partitionContext) {}

// EnterRange_partition_list is called when production range_partition_list is entered.
func (s *BaseDmSqlParserListener) EnterRange_partition_list(ctx *Range_partition_listContext) {}

// ExitRange_partition_list is called when production range_partition_list is exited.
func (s *BaseDmSqlParserListener) ExitRange_partition_list(ctx *Range_partition_listContext) {}

// EnterHash_partition is called when production hash_partition is entered.
func (s *BaseDmSqlParserListener) EnterHash_partition(ctx *Hash_partitionContext) {}

// ExitHash_partition is called when production hash_partition is exited.
func (s *BaseDmSqlParserListener) ExitHash_partition(ctx *Hash_partitionContext) {}

// EnterHash_partition_list is called when production hash_partition_list is entered.
func (s *BaseDmSqlParserListener) EnterHash_partition_list(ctx *Hash_partition_listContext) {}

// ExitHash_partition_list is called when production hash_partition_list is exited.
func (s *BaseDmSqlParserListener) ExitHash_partition_list(ctx *Hash_partition_listContext) {}

// EnterList_partition is called when production list_partition is entered.
func (s *BaseDmSqlParserListener) EnterList_partition(ctx *List_partitionContext) {}

// ExitList_partition is called when production list_partition is exited.
func (s *BaseDmSqlParserListener) ExitList_partition(ctx *List_partitionContext) {}

// EnterList_partition_list is called when production list_partition_list is entered.
func (s *BaseDmSqlParserListener) EnterList_partition_list(ctx *List_partition_listContext) {}

// ExitList_partition_list is called when production list_partition_list is exited.
func (s *BaseDmSqlParserListener) ExitList_partition_list(ctx *List_partition_listContext) {}

// EnterSplit_partition_list is called when production split_partition_list is entered.
func (s *BaseDmSqlParserListener) EnterSplit_partition_list(ctx *Split_partition_listContext) {}

// ExitSplit_partition_list is called when production split_partition_list is exited.
func (s *BaseDmSqlParserListener) ExitSplit_partition_list(ctx *Split_partition_listContext) {}

// EnterPartition_act is called when production partition_act is entered.
func (s *BaseDmSqlParserListener) EnterPartition_act(ctx *Partition_actContext) {}

// ExitPartition_act is called when production partition_act is exited.
func (s *BaseDmSqlParserListener) ExitPartition_act(ctx *Partition_actContext) {}

// EnterVertical_partition_act is called when production vertical_partition_act is entered.
func (s *BaseDmSqlParserListener) EnterVertical_partition_act(ctx *Vertical_partition_actContext) {}

// ExitVertical_partition_act is called when production vertical_partition_act is exited.
func (s *BaseDmSqlParserListener) ExitVertical_partition_act(ctx *Vertical_partition_actContext) {}

// EnterInterval_item is called when production interval_item is entered.
func (s *BaseDmSqlParserListener) EnterInterval_item(ctx *Interval_itemContext) {}

// ExitInterval_item is called when production interval_item is exited.
func (s *BaseDmSqlParserListener) ExitInterval_item(ctx *Interval_itemContext) {}

// EnterHorizon_partition_act_datatype is called when production horizon_partition_act_datatype is entered.
func (s *BaseDmSqlParserListener) EnterHorizon_partition_act_datatype(ctx *Horizon_partition_act_datatypeContext) {
}

// ExitHorizon_partition_act_datatype is called when production horizon_partition_act_datatype is exited.
func (s *BaseDmSqlParserListener) ExitHorizon_partition_act_datatype(ctx *Horizon_partition_act_datatypeContext) {
}

// EnterHorizon_partition_act is called when production horizon_partition_act is entered.
func (s *BaseDmSqlParserListener) EnterHorizon_partition_act(ctx *Horizon_partition_actContext) {}

// ExitHorizon_partition_act is called when production horizon_partition_act is exited.
func (s *BaseDmSqlParserListener) ExitHorizon_partition_act(ctx *Horizon_partition_actContext) {}

// EnterLock_partitions_clause is called when production lock_partitions_clause is entered.
func (s *BaseDmSqlParserListener) EnterLock_partitions_clause(ctx *Lock_partitions_clauseContext) {}

// ExitLock_partitions_clause is called when production lock_partitions_clause is exited.
func (s *BaseDmSqlParserListener) ExitLock_partitions_clause(ctx *Lock_partitions_clauseContext) {}

// EnterSubpartion_template_list_datatype_options is called when production subpartion_template_list_datatype_options is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template_list_datatype_options(ctx *Subpartion_template_list_datatype_optionsContext) {
}

// ExitSubpartion_template_list_datatype_options is called when production subpartion_template_list_datatype_options is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template_list_datatype_options(ctx *Subpartion_template_list_datatype_optionsContext) {
}

// EnterSubpartion_template_list_datatype is called when production subpartion_template_list_datatype is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template_list_datatype(ctx *Subpartion_template_list_datatypeContext) {
}

// ExitSubpartion_template_list_datatype is called when production subpartion_template_list_datatype is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template_list_datatype(ctx *Subpartion_template_list_datatypeContext) {
}

// EnterSubpartion_template_list_options is called when production subpartion_template_list_options is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template_list_options(ctx *Subpartion_template_list_optionsContext) {
}

// ExitSubpartion_template_list_options is called when production subpartion_template_list_options is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template_list_options(ctx *Subpartion_template_list_optionsContext) {
}

// EnterSubpartion_template_list is called when production subpartion_template_list is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template_list(ctx *Subpartion_template_listContext) {
}

// ExitSubpartion_template_list is called when production subpartion_template_list is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template_list(ctx *Subpartion_template_listContext) {
}

// EnterSubpartion_template_datatype is called when production subpartion_template_datatype is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template_datatype(ctx *Subpartion_template_datatypeContext) {
}

// ExitSubpartion_template_datatype is called when production subpartion_template_datatype is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template_datatype(ctx *Subpartion_template_datatypeContext) {
}

// EnterRange_subpartion_template_datatype is called when production range_subpartion_template_datatype is entered.
func (s *BaseDmSqlParserListener) EnterRange_subpartion_template_datatype(ctx *Range_subpartion_template_datatypeContext) {
}

// ExitRange_subpartion_template_datatype is called when production range_subpartion_template_datatype is exited.
func (s *BaseDmSqlParserListener) ExitRange_subpartion_template_datatype(ctx *Range_subpartion_template_datatypeContext) {
}

// EnterHash_subpartion_template_datatype is called when production hash_subpartion_template_datatype is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartion_template_datatype(ctx *Hash_subpartion_template_datatypeContext) {
}

// ExitHash_subpartion_template_datatype is called when production hash_subpartion_template_datatype is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartion_template_datatype(ctx *Hash_subpartion_template_datatypeContext) {
}

// EnterHash_subpartions_template_datatype_option is called when production hash_subpartions_template_datatype_option is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartions_template_datatype_option(ctx *Hash_subpartions_template_datatype_optionContext) {
}

// ExitHash_subpartions_template_datatype_option is called when production hash_subpartions_template_datatype_option is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartions_template_datatype_option(ctx *Hash_subpartions_template_datatype_optionContext) {
}

// EnterList_subpartion_template_datatype is called when production list_subpartion_template_datatype is entered.
func (s *BaseDmSqlParserListener) EnterList_subpartion_template_datatype(ctx *List_subpartion_template_datatypeContext) {
}

// ExitList_subpartion_template_datatype is called when production list_subpartion_template_datatype is exited.
func (s *BaseDmSqlParserListener) ExitList_subpartion_template_datatype(ctx *List_subpartion_template_datatypeContext) {
}

// EnterSubpartion_template is called when production subpartion_template is entered.
func (s *BaseDmSqlParserListener) EnterSubpartion_template(ctx *Subpartion_templateContext) {}

// ExitSubpartion_template is called when production subpartion_template is exited.
func (s *BaseDmSqlParserListener) ExitSubpartion_template(ctx *Subpartion_templateContext) {}

// EnterRange_subpartion_template is called when production range_subpartion_template is entered.
func (s *BaseDmSqlParserListener) EnterRange_subpartion_template(ctx *Range_subpartion_templateContext) {
}

// ExitRange_subpartion_template is called when production range_subpartion_template is exited.
func (s *BaseDmSqlParserListener) ExitRange_subpartion_template(ctx *Range_subpartion_templateContext) {
}

// EnterHash_subpartion_template is called when production hash_subpartion_template is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartion_template(ctx *Hash_subpartion_templateContext) {
}

// ExitHash_subpartion_template is called when production hash_subpartion_template is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartion_template(ctx *Hash_subpartion_templateContext) {
}

// EnterHash_subpartions_template_option is called when production hash_subpartions_template_option is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartions_template_option(ctx *Hash_subpartions_template_optionContext) {
}

// ExitHash_subpartions_template_option is called when production hash_subpartions_template_option is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartions_template_option(ctx *Hash_subpartions_template_optionContext) {
}

// EnterList_subpartion_template is called when production list_subpartion_template is entered.
func (s *BaseDmSqlParserListener) EnterList_subpartion_template(ctx *List_subpartion_templateContext) {
}

// ExitList_subpartion_template is called when production list_subpartion_template is exited.
func (s *BaseDmSqlParserListener) ExitList_subpartion_template(ctx *List_subpartion_templateContext) {
}

// EnterRange_subpartition is called when production range_subpartition is entered.
func (s *BaseDmSqlParserListener) EnterRange_subpartition(ctx *Range_subpartitionContext) {}

// ExitRange_subpartition is called when production range_subpartition is exited.
func (s *BaseDmSqlParserListener) ExitRange_subpartition(ctx *Range_subpartitionContext) {}

// EnterHash_subpartition is called when production hash_subpartition is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartition(ctx *Hash_subpartitionContext) {}

// ExitHash_subpartition is called when production hash_subpartition is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartition(ctx *Hash_subpartitionContext) {}

// EnterList_subpartition is called when production list_subpartition is entered.
func (s *BaseDmSqlParserListener) EnterList_subpartition(ctx *List_subpartitionContext) {}

// ExitList_subpartition is called when production list_subpartition is exited.
func (s *BaseDmSqlParserListener) ExitList_subpartition(ctx *List_subpartitionContext) {}

// EnterRange_subpartition_list is called when production range_subpartition_list is entered.
func (s *BaseDmSqlParserListener) EnterRange_subpartition_list(ctx *Range_subpartition_listContext) {}

// ExitRange_subpartition_list is called when production range_subpartition_list is exited.
func (s *BaseDmSqlParserListener) ExitRange_subpartition_list(ctx *Range_subpartition_listContext) {}

// EnterHash_subpartition_list is called when production hash_subpartition_list is entered.
func (s *BaseDmSqlParserListener) EnterHash_subpartition_list(ctx *Hash_subpartition_listContext) {}

// ExitHash_subpartition_list is called when production hash_subpartition_list is exited.
func (s *BaseDmSqlParserListener) ExitHash_subpartition_list(ctx *Hash_subpartition_listContext) {}

// EnterList_subpartition_list is called when production list_subpartition_list is entered.
func (s *BaseDmSqlParserListener) EnterList_subpartition_list(ctx *List_subpartition_listContext) {}

// ExitList_subpartition_list is called when production list_subpartition_list is exited.
func (s *BaseDmSqlParserListener) ExitList_subpartition_list(ctx *List_subpartition_listContext) {}

// EnterSubpartition_hash_desc is called when production subpartition_hash_desc is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_hash_desc(ctx *Subpartition_hash_descContext) {}

// ExitSubpartition_hash_desc is called when production subpartition_hash_desc is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_hash_desc(ctx *Subpartition_hash_descContext) {}

// EnterSubpartition_range_desc is called when production subpartition_range_desc is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_range_desc(ctx *Subpartition_range_descContext) {}

// ExitSubpartition_range_desc is called when production subpartition_range_desc is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_range_desc(ctx *Subpartition_range_descContext) {}

// EnterSubpartition_list_desc is called when production subpartition_list_desc is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_list_desc(ctx *Subpartition_list_descContext) {}

// ExitSubpartition_list_desc is called when production subpartition_list_desc is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_list_desc(ctx *Subpartition_list_descContext) {}

// EnterSubpartition_hash_desc_list is called when production subpartition_hash_desc_list is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_hash_desc_list(ctx *Subpartition_hash_desc_listContext) {
}

// ExitSubpartition_hash_desc_list is called when production subpartition_hash_desc_list is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_hash_desc_list(ctx *Subpartition_hash_desc_listContext) {
}

// EnterSubpartition_range_desc_list is called when production subpartition_range_desc_list is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_range_desc_list(ctx *Subpartition_range_desc_listContext) {
}

// ExitSubpartition_range_desc_list is called when production subpartition_range_desc_list is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_range_desc_list(ctx *Subpartition_range_desc_listContext) {
}

// EnterSubpartition_list_desc_list is called when production subpartition_list_desc_list is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_list_desc_list(ctx *Subpartition_list_desc_listContext) {
}

// ExitSubpartition_list_desc_list is called when production subpartition_list_desc_list is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_list_desc_list(ctx *Subpartition_list_desc_listContext) {
}

// EnterSubpartition_desc is called when production subpartition_desc is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_desc(ctx *Subpartition_descContext) {}

// ExitSubpartition_desc is called when production subpartition_desc is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_desc(ctx *Subpartition_descContext) {}

// EnterSubpartition_desc_option is called when production subpartition_desc_option is entered.
func (s *BaseDmSqlParserListener) EnterSubpartition_desc_option(ctx *Subpartition_desc_optionContext) {
}

// ExitSubpartition_desc_option is called when production subpartition_desc_option is exited.
func (s *BaseDmSqlParserListener) ExitSubpartition_desc_option(ctx *Subpartition_desc_optionContext) {
}

// EnterAdd_subpartition_desc is called when production add_subpartition_desc is entered.
func (s *BaseDmSqlParserListener) EnterAdd_subpartition_desc(ctx *Add_subpartition_descContext) {}

// ExitAdd_subpartition_desc is called when production add_subpartition_desc is exited.
func (s *BaseDmSqlParserListener) ExitAdd_subpartition_desc(ctx *Add_subpartition_descContext) {}

// EnterPartition_no is called when production partition_no is entered.
func (s *BaseDmSqlParserListener) EnterPartition_no(ctx *Partition_noContext) {}

// ExitPartition_no is called when production partition_no is exited.
func (s *BaseDmSqlParserListener) ExitPartition_no(ctx *Partition_noContext) {}

// EnterComment_clause is called when production comment_clause is entered.
func (s *BaseDmSqlParserListener) EnterComment_clause(ctx *Comment_clauseContext) {}

// ExitComment_clause is called when production comment_clause is exited.
func (s *BaseDmSqlParserListener) ExitComment_clause(ctx *Comment_clauseContext) {}

// EnterEncrypt_clause_options is called when production encrypt_clause_options is entered.
func (s *BaseDmSqlParserListener) EnterEncrypt_clause_options(ctx *Encrypt_clause_optionsContext) {}

// ExitEncrypt_clause_options is called when production encrypt_clause_options is exited.
func (s *BaseDmSqlParserListener) ExitEncrypt_clause_options(ctx *Encrypt_clause_optionsContext) {}

// EnterEncrypt_clause is called when production encrypt_clause is entered.
func (s *BaseDmSqlParserListener) EnterEncrypt_clause(ctx *Encrypt_clauseContext) {}

// ExitEncrypt_clause is called when production encrypt_clause is exited.
func (s *BaseDmSqlParserListener) ExitEncrypt_clause(ctx *Encrypt_clauseContext) {}

// EnterEncrypt_cipher is called when production encrypt_cipher is entered.
func (s *BaseDmSqlParserListener) EnterEncrypt_cipher(ctx *Encrypt_cipherContext) {}

// ExitEncrypt_cipher is called when production encrypt_cipher is exited.
func (s *BaseDmSqlParserListener) ExitEncrypt_cipher(ctx *Encrypt_cipherContext) {}

// EnterCrypto_name is called when production crypto_name is entered.
func (s *BaseDmSqlParserListener) EnterCrypto_name(ctx *Crypto_nameContext) {}

// ExitCrypto_name is called when production crypto_name is exited.
func (s *BaseDmSqlParserListener) ExitCrypto_name(ctx *Crypto_nameContext) {}

// EnterCipher_name is called when production cipher_name is entered.
func (s *BaseDmSqlParserListener) EnterCipher_name(ctx *Cipher_nameContext) {}

// ExitCipher_name is called when production cipher_name is exited.
func (s *BaseDmSqlParserListener) ExitCipher_name(ctx *Cipher_nameContext) {}

// EnterFull_cipher_name is called when production full_cipher_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_cipher_name(ctx *Full_cipher_nameContext) {}

// ExitFull_cipher_name is called when production full_cipher_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_cipher_name(ctx *Full_cipher_nameContext) {}

// EnterEncrypt_type is called when production encrypt_type is entered.
func (s *BaseDmSqlParserListener) EnterEncrypt_type(ctx *Encrypt_typeContext) {}

// ExitEncrypt_type is called when production encrypt_type is exited.
func (s *BaseDmSqlParserListener) ExitEncrypt_type(ctx *Encrypt_typeContext) {}

// EnterManual_clause is called when production manual_clause is entered.
func (s *BaseDmSqlParserListener) EnterManual_clause(ctx *Manual_clauseContext) {}

// ExitManual_clause is called when production manual_clause is exited.
func (s *BaseDmSqlParserListener) ExitManual_clause(ctx *Manual_clauseContext) {}

// EnterUser_clause_options is called when production user_clause_options is entered.
func (s *BaseDmSqlParserListener) EnterUser_clause_options(ctx *User_clause_optionsContext) {}

// ExitUser_clause_options is called when production user_clause_options is exited.
func (s *BaseDmSqlParserListener) ExitUser_clause_options(ctx *User_clause_optionsContext) {}

// EnterUser_clause is called when production user_clause is entered.
func (s *BaseDmSqlParserListener) EnterUser_clause(ctx *User_clauseContext) {}

// ExitUser_clause is called when production user_clause is exited.
func (s *BaseDmSqlParserListener) ExitUser_clause(ctx *User_clauseContext) {}

// EnterUser_list_option is called when production user_list_option is entered.
func (s *BaseDmSqlParserListener) EnterUser_list_option(ctx *User_list_optionContext) {}

// ExitUser_list_option is called when production user_list_option is exited.
func (s *BaseDmSqlParserListener) ExitUser_list_option(ctx *User_list_optionContext) {}

// EnterUser_list is called when production user_list is entered.
func (s *BaseDmSqlParserListener) EnterUser_list(ctx *User_listContext) {}

// ExitUser_list is called when production user_list is exited.
func (s *BaseDmSqlParserListener) ExitUser_list(ctx *User_listContext) {}

// EnterHash_cipher is called when production hash_cipher is entered.
func (s *BaseDmSqlParserListener) EnterHash_cipher(ctx *Hash_cipherContext) {}

// ExitHash_cipher is called when production hash_cipher is exited.
func (s *BaseDmSqlParserListener) ExitHash_cipher(ctx *Hash_cipherContext) {}

// EnterHash_type is called when production hash_type is entered.
func (s *BaseDmSqlParserListener) EnterHash_type(ctx *Hash_typeContext) {}

// ExitHash_type is called when production hash_type is exited.
func (s *BaseDmSqlParserListener) ExitHash_type(ctx *Hash_typeContext) {}

// EnterSpace_limit is called when production space_limit is entered.
func (s *BaseDmSqlParserListener) EnterSpace_limit(ctx *Space_limitContext) {}

// ExitSpace_limit is called when production space_limit is exited.
func (s *BaseDmSqlParserListener) ExitSpace_limit(ctx *Space_limitContext) {}

// EnterSpace_limit_nn is called when production space_limit_nn is entered.
func (s *BaseDmSqlParserListener) EnterSpace_limit_nn(ctx *Space_limit_nnContext) {}

// ExitSpace_limit_nn is called when production space_limit_nn is exited.
func (s *BaseDmSqlParserListener) ExitSpace_limit_nn(ctx *Space_limit_nnContext) {}

// EnterSpace_limit_1 is called when production space_limit_1 is entered.
func (s *BaseDmSqlParserListener) EnterSpace_limit_1(ctx *Space_limit_1Context) {}

// ExitSpace_limit_1 is called when production space_limit_1 is exited.
func (s *BaseDmSqlParserListener) ExitSpace_limit_1(ctx *Space_limit_1Context) {}

// EnterSpace_limit2 is called when production space_limit2 is entered.
func (s *BaseDmSqlParserListener) EnterSpace_limit2(ctx *Space_limit2Context) {}

// ExitSpace_limit2 is called when production space_limit2 is exited.
func (s *BaseDmSqlParserListener) ExitSpace_limit2(ctx *Space_limit2Context) {}

// EnterDel_res is called when production del_res is entered.
func (s *BaseDmSqlParserListener) EnterDel_res(ctx *Del_resContext) {}

// ExitDel_res is called when production del_res is exited.
func (s *BaseDmSqlParserListener) ExitDel_res(ctx *Del_resContext) {}

// EnterTrig_enable is called when production trig_enable is entered.
func (s *BaseDmSqlParserListener) EnterTrig_enable(ctx *Trig_enableContext) {}

// ExitTrig_enable is called when production trig_enable is exited.
func (s *BaseDmSqlParserListener) ExitTrig_enable(ctx *Trig_enableContext) {}

// EnterAt_raft is called when production at_raft is entered.
func (s *BaseDmSqlParserListener) EnterAt_raft(ctx *At_raftContext) {}

// ExitAt_raft is called when production at_raft is exited.
func (s *BaseDmSqlParserListener) ExitAt_raft(ctx *At_raftContext) {}

// EnterCreate_trigger_stmt is called when production create_trigger_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// ExitCreate_trigger_stmt is called when production create_trigger_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_trigger_stmt(ctx *Create_trigger_stmtContext) {}

// EnterBefore_after is called when production before_after is entered.
func (s *BaseDmSqlParserListener) EnterBefore_after(ctx *Before_afterContext) {}

// ExitBefore_after is called when production before_after is exited.
func (s *BaseDmSqlParserListener) ExitBefore_after(ctx *Before_afterContext) {}

// EnterTrig_del_ins_upd_list is called when production trig_del_ins_upd_list is entered.
func (s *BaseDmSqlParserListener) EnterTrig_del_ins_upd_list(ctx *Trig_del_ins_upd_listContext) {}

// ExitTrig_del_ins_upd_list is called when production trig_del_ins_upd_list is exited.
func (s *BaseDmSqlParserListener) ExitTrig_del_ins_upd_list(ctx *Trig_del_ins_upd_listContext) {}

// EnterTrig_del_ins_upd is called when production trig_del_ins_upd is entered.
func (s *BaseDmSqlParserListener) EnterTrig_del_ins_upd(ctx *Trig_del_ins_updContext) {}

// ExitTrig_del_ins_upd is called when production trig_del_ins_upd is exited.
func (s *BaseDmSqlParserListener) ExitTrig_del_ins_upd(ctx *Trig_del_ins_updContext) {}

// EnterUpdate_of_option is called when production update_of_option is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_of_option(ctx *Update_of_optionContext) {}

// ExitUpdate_of_option is called when production update_of_option is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_of_option(ctx *Update_of_optionContext) {}

// EnterNowait is called when production nowait is entered.
func (s *BaseDmSqlParserListener) EnterNowait(ctx *NowaitContext) {}

// ExitNowait is called when production nowait is exited.
func (s *BaseDmSqlParserListener) ExitNowait(ctx *NowaitContext) {}

// EnterTrig_event_list is called when production trig_event_list is entered.
func (s *BaseDmSqlParserListener) EnterTrig_event_list(ctx *Trig_event_listContext) {}

// ExitTrig_event_list is called when production trig_event_list is exited.
func (s *BaseDmSqlParserListener) ExitTrig_event_list(ctx *Trig_event_listContext) {}

// EnterTrig_event is called when production trig_event is entered.
func (s *BaseDmSqlParserListener) EnterTrig_event(ctx *Trig_eventContext) {}

// ExitTrig_event is called when production trig_event is exited.
func (s *BaseDmSqlParserListener) ExitTrig_event(ctx *Trig_eventContext) {}

// EnterEvent_object is called when production event_object is entered.
func (s *BaseDmSqlParserListener) EnterEvent_object(ctx *Event_objectContext) {}

// ExitEvent_object is called when production event_object is exited.
func (s *BaseDmSqlParserListener) ExitEvent_object(ctx *Event_objectContext) {}

// EnterTrig_referencing_def_options is called when production trig_referencing_def_options is entered.
func (s *BaseDmSqlParserListener) EnterTrig_referencing_def_options(ctx *Trig_referencing_def_optionsContext) {
}

// ExitTrig_referencing_def_options is called when production trig_referencing_def_options is exited.
func (s *BaseDmSqlParserListener) ExitTrig_referencing_def_options(ctx *Trig_referencing_def_optionsContext) {
}

// EnterTrig_referencing_def is called when production trig_referencing_def is entered.
func (s *BaseDmSqlParserListener) EnterTrig_referencing_def(ctx *Trig_referencing_defContext) {}

// ExitTrig_referencing_def is called when production trig_referencing_def is exited.
func (s *BaseDmSqlParserListener) ExitTrig_referencing_def(ctx *Trig_referencing_defContext) {}

// EnterTrig_referencing_list is called when production trig_referencing_list is entered.
func (s *BaseDmSqlParserListener) EnterTrig_referencing_list(ctx *Trig_referencing_listContext) {}

// ExitTrig_referencing_list is called when production trig_referencing_list is exited.
func (s *BaseDmSqlParserListener) ExitTrig_referencing_list(ctx *Trig_referencing_listContext) {}

// EnterTrig_referencing_old is called when production trig_referencing_old is entered.
func (s *BaseDmSqlParserListener) EnterTrig_referencing_old(ctx *Trig_referencing_oldContext) {}

// ExitTrig_referencing_old is called when production trig_referencing_old is exited.
func (s *BaseDmSqlParserListener) ExitTrig_referencing_old(ctx *Trig_referencing_oldContext) {}

// EnterTrig_referencing_new is called when production trig_referencing_new is entered.
func (s *BaseDmSqlParserListener) EnterTrig_referencing_new(ctx *Trig_referencing_newContext) {}

// ExitTrig_referencing_new is called when production trig_referencing_new is exited.
func (s *BaseDmSqlParserListener) ExitTrig_referencing_new(ctx *Trig_referencing_newContext) {}

// EnterTrig_for_each_option is called when production trig_for_each_option is entered.
func (s *BaseDmSqlParserListener) EnterTrig_for_each_option(ctx *Trig_for_each_optionContext) {}

// ExitTrig_for_each_option is called when production trig_for_each_option is exited.
func (s *BaseDmSqlParserListener) ExitTrig_for_each_option(ctx *Trig_for_each_optionContext) {}

// EnterTrig_timer_rate is called when production trig_timer_rate is entered.
func (s *BaseDmSqlParserListener) EnterTrig_timer_rate(ctx *Trig_timer_rateContext) {}

// ExitTrig_timer_rate is called when production trig_timer_rate is exited.
func (s *BaseDmSqlParserListener) ExitTrig_timer_rate(ctx *Trig_timer_rateContext) {}

// EnterExec_ep_seqno is called when production exec_ep_seqno is entered.
func (s *BaseDmSqlParserListener) EnterExec_ep_seqno(ctx *Exec_ep_seqnoContext) {}

// ExitExec_ep_seqno is called when production exec_ep_seqno is exited.
func (s *BaseDmSqlParserListener) ExitExec_ep_seqno(ctx *Exec_ep_seqnoContext) {}

// EnterRate_over_day is called when production rate_over_day is entered.
func (s *BaseDmSqlParserListener) EnterRate_over_day(ctx *Rate_over_dayContext) {}

// ExitRate_over_day is called when production rate_over_day is exited.
func (s *BaseDmSqlParserListener) ExitRate_over_day(ctx *Rate_over_dayContext) {}

// EnterMonth_rate is called when production month_rate is entered.
func (s *BaseDmSqlParserListener) EnterMonth_rate(ctx *Month_rateContext) {}

// ExitMonth_rate is called when production month_rate is exited.
func (s *BaseDmSqlParserListener) ExitMonth_rate(ctx *Month_rateContext) {}

// EnterDay_in_month is called when production day_in_month is entered.
func (s *BaseDmSqlParserListener) EnterDay_in_month(ctx *Day_in_monthContext) {}

// ExitDay_in_month is called when production day_in_month is exited.
func (s *BaseDmSqlParserListener) ExitDay_in_month(ctx *Day_in_monthContext) {}

// EnterDay_in_month_week is called when production day_in_month_week is entered.
func (s *BaseDmSqlParserListener) EnterDay_in_month_week(ctx *Day_in_month_weekContext) {}

// ExitDay_in_month_week is called when production day_in_month_week is exited.
func (s *BaseDmSqlParserListener) ExitDay_in_month_week(ctx *Day_in_month_weekContext) {}

// EnterWeek_rate is called when production week_rate is entered.
func (s *BaseDmSqlParserListener) EnterWeek_rate(ctx *Week_rateContext) {}

// ExitWeek_rate is called when production week_rate is exited.
func (s *BaseDmSqlParserListener) ExitWeek_rate(ctx *Week_rateContext) {}

// EnterDay_of_week_list is called when production day_of_week_list is entered.
func (s *BaseDmSqlParserListener) EnterDay_of_week_list(ctx *Day_of_week_listContext) {}

// ExitDay_of_week_list is called when production day_of_week_list is exited.
func (s *BaseDmSqlParserListener) ExitDay_of_week_list(ctx *Day_of_week_listContext) {}

// EnterDay_rate is called when production day_rate is entered.
func (s *BaseDmSqlParserListener) EnterDay_rate(ctx *Day_rateContext) {}

// ExitDay_rate is called when production day_rate is exited.
func (s *BaseDmSqlParserListener) ExitDay_rate(ctx *Day_rateContext) {}

// EnterRate_in_day is called when production rate_in_day is entered.
func (s *BaseDmSqlParserListener) EnterRate_in_day(ctx *Rate_in_dayContext) {}

// ExitRate_in_day is called when production rate_in_day is exited.
func (s *BaseDmSqlParserListener) ExitRate_in_day(ctx *Rate_in_dayContext) {}

// EnterOnce_in_day is called when production once_in_day is entered.
func (s *BaseDmSqlParserListener) EnterOnce_in_day(ctx *Once_in_dayContext) {}

// ExitOnce_in_day is called when production once_in_day is exited.
func (s *BaseDmSqlParserListener) ExitOnce_in_day(ctx *Once_in_dayContext) {}

// EnterTimes_in_day is called when production times_in_day is entered.
func (s *BaseDmSqlParserListener) EnterTimes_in_day(ctx *Times_in_dayContext) {}

// ExitTimes_in_day is called when production times_in_day is exited.
func (s *BaseDmSqlParserListener) ExitTimes_in_day(ctx *Times_in_dayContext) {}

// EnterDuaring_time is called when production duaring_time is entered.
func (s *BaseDmSqlParserListener) EnterDuaring_time(ctx *Duaring_timeContext) {}

// ExitDuaring_time is called when production duaring_time is exited.
func (s *BaseDmSqlParserListener) ExitDuaring_time(ctx *Duaring_timeContext) {}

// EnterDuaring_date is called when production duaring_date is entered.
func (s *BaseDmSqlParserListener) EnterDuaring_date(ctx *Duaring_dateContext) {}

// ExitDuaring_date is called when production duaring_date is exited.
func (s *BaseDmSqlParserListener) ExitDuaring_date(ctx *Duaring_dateContext) {}

// EnterTrig_when_option is called when production trig_when_option is entered.
func (s *BaseDmSqlParserListener) EnterTrig_when_option(ctx *Trig_when_optionContext) {}

// ExitTrig_when_option is called when production trig_when_option is exited.
func (s *BaseDmSqlParserListener) ExitTrig_when_option(ctx *Trig_when_optionContext) {}

// EnterTrig_when_condition is called when production trig_when_condition is entered.
func (s *BaseDmSqlParserListener) EnterTrig_when_condition(ctx *Trig_when_conditionContext) {}

// ExitTrig_when_condition is called when production trig_when_condition is exited.
func (s *BaseDmSqlParserListener) ExitTrig_when_condition(ctx *Trig_when_conditionContext) {}

// EnterRepeat_interval_stmt is called when production repeat_interval_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRepeat_interval_stmt(ctx *Repeat_interval_stmtContext) {}

// ExitRepeat_interval_stmt is called when production repeat_interval_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRepeat_interval_stmt(ctx *Repeat_interval_stmtContext) {}

// EnterMax_run_duration is called when production max_run_duration is entered.
func (s *BaseDmSqlParserListener) EnterMax_run_duration(ctx *Max_run_durationContext) {}

// ExitMax_run_duration is called when production max_run_duration is exited.
func (s *BaseDmSqlParserListener) ExitMax_run_duration(ctx *Max_run_durationContext) {}

// EnterRepeat_interval is called when production repeat_interval is entered.
func (s *BaseDmSqlParserListener) EnterRepeat_interval(ctx *Repeat_intervalContext) {}

// ExitRepeat_interval is called when production repeat_interval is exited.
func (s *BaseDmSqlParserListener) ExitRepeat_interval(ctx *Repeat_intervalContext) {}

// EnterFrequency_clause is called when production frequency_clause is entered.
func (s *BaseDmSqlParserListener) EnterFrequency_clause(ctx *Frequency_clauseContext) {}

// ExitFrequency_clause is called when production frequency_clause is exited.
func (s *BaseDmSqlParserListener) ExitFrequency_clause(ctx *Frequency_clauseContext) {}

// EnterFrequency_define is called when production frequency_define is entered.
func (s *BaseDmSqlParserListener) EnterFrequency_define(ctx *Frequency_defineContext) {}

// ExitFrequency_define is called when production frequency_define is exited.
func (s *BaseDmSqlParserListener) ExitFrequency_define(ctx *Frequency_defineContext) {}

// EnterPredefined_frequency is called when production predefined_frequency is entered.
func (s *BaseDmSqlParserListener) EnterPredefined_frequency(ctx *Predefined_frequencyContext) {}

// ExitPredefined_frequency is called when production predefined_frequency is exited.
func (s *BaseDmSqlParserListener) ExitPredefined_frequency(ctx *Predefined_frequencyContext) {}

// EnterInterval_clause_list is called when production interval_clause_list is entered.
func (s *BaseDmSqlParserListener) EnterInterval_clause_list(ctx *Interval_clause_listContext) {}

// ExitInterval_clause_list is called when production interval_clause_list is exited.
func (s *BaseDmSqlParserListener) ExitInterval_clause_list(ctx *Interval_clause_listContext) {}

// EnterInterval_clause_single is called when production interval_clause_single is entered.
func (s *BaseDmSqlParserListener) EnterInterval_clause_single(ctx *Interval_clause_singleContext) {}

// ExitInterval_clause_single is called when production interval_clause_single is exited.
func (s *BaseDmSqlParserListener) ExitInterval_clause_single(ctx *Interval_clause_singleContext) {}

// EnterInterval_clause is called when production interval_clause is entered.
func (s *BaseDmSqlParserListener) EnterInterval_clause(ctx *Interval_clauseContext) {}

// ExitInterval_clause is called when production interval_clause is exited.
func (s *BaseDmSqlParserListener) ExitInterval_clause(ctx *Interval_clauseContext) {}

// EnterIntervalnum is called when production intervalnum is entered.
func (s *BaseDmSqlParserListener) EnterIntervalnum(ctx *IntervalnumContext) {}

// ExitIntervalnum is called when production intervalnum is exited.
func (s *BaseDmSqlParserListener) ExitIntervalnum(ctx *IntervalnumContext) {}

// EnterBymonth_clause is called when production bymonth_clause is entered.
func (s *BaseDmSqlParserListener) EnterBymonth_clause(ctx *Bymonth_clauseContext) {}

// ExitBymonth_clause is called when production bymonth_clause is exited.
func (s *BaseDmSqlParserListener) ExitBymonth_clause(ctx *Bymonth_clauseContext) {}

// EnterMonthlist is called when production monthlist is entered.
func (s *BaseDmSqlParserListener) EnterMonthlist(ctx *MonthlistContext) {}

// ExitMonthlist is called when production monthlist is exited.
func (s *BaseDmSqlParserListener) ExitMonthlist(ctx *MonthlistContext) {}

// EnterMonth is called when production month is entered.
func (s *BaseDmSqlParserListener) EnterMonth(ctx *MonthContext) {}

// ExitMonth is called when production month is exited.
func (s *BaseDmSqlParserListener) ExitMonth(ctx *MonthContext) {}

// EnterNumeric_month is called when production numeric_month is entered.
func (s *BaseDmSqlParserListener) EnterNumeric_month(ctx *Numeric_monthContext) {}

// ExitNumeric_month is called when production numeric_month is exited.
func (s *BaseDmSqlParserListener) ExitNumeric_month(ctx *Numeric_monthContext) {}

// EnterChar_month is called when production char_month is entered.
func (s *BaseDmSqlParserListener) EnterChar_month(ctx *Char_monthContext) {}

// ExitChar_month is called when production char_month is exited.
func (s *BaseDmSqlParserListener) ExitChar_month(ctx *Char_monthContext) {}

// EnterByweekno_clause is called when production byweekno_clause is entered.
func (s *BaseDmSqlParserListener) EnterByweekno_clause(ctx *Byweekno_clauseContext) {}

// ExitByweekno_clause is called when production byweekno_clause is exited.
func (s *BaseDmSqlParserListener) ExitByweekno_clause(ctx *Byweekno_clauseContext) {}

// EnterWeekno_list is called when production weekno_list is entered.
func (s *BaseDmSqlParserListener) EnterWeekno_list(ctx *Weekno_listContext) {}

// ExitWeekno_list is called when production weekno_list is exited.
func (s *BaseDmSqlParserListener) ExitWeekno_list(ctx *Weekno_listContext) {}

// EnterWeekno is called when production weekno is entered.
func (s *BaseDmSqlParserListener) EnterWeekno(ctx *WeeknoContext) {}

// ExitWeekno is called when production weekno is exited.
func (s *BaseDmSqlParserListener) ExitWeekno(ctx *WeeknoContext) {}

// EnterByyearday_clause is called when production byyearday_clause is entered.
func (s *BaseDmSqlParserListener) EnterByyearday_clause(ctx *Byyearday_clauseContext) {}

// ExitByyearday_clause is called when production byyearday_clause is exited.
func (s *BaseDmSqlParserListener) ExitByyearday_clause(ctx *Byyearday_clauseContext) {}

// EnterYearday_list is called when production yearday_list is entered.
func (s *BaseDmSqlParserListener) EnterYearday_list(ctx *Yearday_listContext) {}

// ExitYearday_list is called when production yearday_list is exited.
func (s *BaseDmSqlParserListener) ExitYearday_list(ctx *Yearday_listContext) {}

// EnterYearday is called when production yearday is entered.
func (s *BaseDmSqlParserListener) EnterYearday(ctx *YeardayContext) {}

// ExitYearday is called when production yearday is exited.
func (s *BaseDmSqlParserListener) ExitYearday(ctx *YeardayContext) {}

// EnterBymonthday_clause is called when production bymonthday_clause is entered.
func (s *BaseDmSqlParserListener) EnterBymonthday_clause(ctx *Bymonthday_clauseContext) {}

// ExitBymonthday_clause is called when production bymonthday_clause is exited.
func (s *BaseDmSqlParserListener) ExitBymonthday_clause(ctx *Bymonthday_clauseContext) {}

// EnterMonthday_list is called when production monthday_list is entered.
func (s *BaseDmSqlParserListener) EnterMonthday_list(ctx *Monthday_listContext) {}

// ExitMonthday_list is called when production monthday_list is exited.
func (s *BaseDmSqlParserListener) ExitMonthday_list(ctx *Monthday_listContext) {}

// EnterMonthday is called when production monthday is entered.
func (s *BaseDmSqlParserListener) EnterMonthday(ctx *MonthdayContext) {}

// ExitMonthday is called when production monthday is exited.
func (s *BaseDmSqlParserListener) ExitMonthday(ctx *MonthdayContext) {}

// EnterByday_clause is called when production byday_clause is entered.
func (s *BaseDmSqlParserListener) EnterByday_clause(ctx *Byday_clauseContext) {}

// ExitByday_clause is called when production byday_clause is exited.
func (s *BaseDmSqlParserListener) ExitByday_clause(ctx *Byday_clauseContext) {}

// EnterByday_list is called when production byday_list is entered.
func (s *BaseDmSqlParserListener) EnterByday_list(ctx *Byday_listContext) {}

// ExitByday_list is called when production byday_list is exited.
func (s *BaseDmSqlParserListener) ExitByday_list(ctx *Byday_listContext) {}

// EnterByday is called when production byday is entered.
func (s *BaseDmSqlParserListener) EnterByday(ctx *BydayContext) {}

// ExitByday is called when production byday is exited.
func (s *BaseDmSqlParserListener) ExitByday(ctx *BydayContext) {}

// EnterWeekdaynum_options is called when production weekdaynum_options is entered.
func (s *BaseDmSqlParserListener) EnterWeekdaynum_options(ctx *Weekdaynum_optionsContext) {}

// ExitWeekdaynum_options is called when production weekdaynum_options is exited.
func (s *BaseDmSqlParserListener) ExitWeekdaynum_options(ctx *Weekdaynum_optionsContext) {}

// EnterWeekdaynum is called when production weekdaynum is entered.
func (s *BaseDmSqlParserListener) EnterWeekdaynum(ctx *WeekdaynumContext) {}

// ExitWeekdaynum is called when production weekdaynum is exited.
func (s *BaseDmSqlParserListener) ExitWeekdaynum(ctx *WeekdaynumContext) {}

// EnterDay is called when production day is entered.
func (s *BaseDmSqlParserListener) EnterDay(ctx *DayContext) {}

// ExitDay is called when production day is exited.
func (s *BaseDmSqlParserListener) ExitDay(ctx *DayContext) {}

// EnterByhour_clause is called when production byhour_clause is entered.
func (s *BaseDmSqlParserListener) EnterByhour_clause(ctx *Byhour_clauseContext) {}

// ExitByhour_clause is called when production byhour_clause is exited.
func (s *BaseDmSqlParserListener) ExitByhour_clause(ctx *Byhour_clauseContext) {}

// EnterHour_list is called when production hour_list is entered.
func (s *BaseDmSqlParserListener) EnterHour_list(ctx *Hour_listContext) {}

// ExitHour_list is called when production hour_list is exited.
func (s *BaseDmSqlParserListener) ExitHour_list(ctx *Hour_listContext) {}

// EnterHour is called when production hour is entered.
func (s *BaseDmSqlParserListener) EnterHour(ctx *HourContext) {}

// ExitHour is called when production hour is exited.
func (s *BaseDmSqlParserListener) ExitHour(ctx *HourContext) {}

// EnterByminute_clause is called when production byminute_clause is entered.
func (s *BaseDmSqlParserListener) EnterByminute_clause(ctx *Byminute_clauseContext) {}

// ExitByminute_clause is called when production byminute_clause is exited.
func (s *BaseDmSqlParserListener) ExitByminute_clause(ctx *Byminute_clauseContext) {}

// EnterMinute_list is called when production minute_list is entered.
func (s *BaseDmSqlParserListener) EnterMinute_list(ctx *Minute_listContext) {}

// ExitMinute_list is called when production minute_list is exited.
func (s *BaseDmSqlParserListener) ExitMinute_list(ctx *Minute_listContext) {}

// EnterMinute is called when production minute is entered.
func (s *BaseDmSqlParserListener) EnterMinute(ctx *MinuteContext) {}

// ExitMinute is called when production minute is exited.
func (s *BaseDmSqlParserListener) ExitMinute(ctx *MinuteContext) {}

// EnterBysecond_clause is called when production bysecond_clause is entered.
func (s *BaseDmSqlParserListener) EnterBysecond_clause(ctx *Bysecond_clauseContext) {}

// ExitBysecond_clause is called when production bysecond_clause is exited.
func (s *BaseDmSqlParserListener) ExitBysecond_clause(ctx *Bysecond_clauseContext) {}

// EnterSecond_list is called when production second_list is entered.
func (s *BaseDmSqlParserListener) EnterSecond_list(ctx *Second_listContext) {}

// ExitSecond_list is called when production second_list is exited.
func (s *BaseDmSqlParserListener) ExitSecond_list(ctx *Second_listContext) {}

// EnterSecond is called when production second is entered.
func (s *BaseDmSqlParserListener) EnterSecond(ctx *SecondContext) {}

// ExitSecond is called when production second is exited.
func (s *BaseDmSqlParserListener) ExitSecond(ctx *SecondContext) {}

// EnterQuery_rewrite is called when production query_rewrite is entered.
func (s *BaseDmSqlParserListener) EnterQuery_rewrite(ctx *Query_rewriteContext) {}

// ExitQuery_rewrite is called when production query_rewrite is exited.
func (s *BaseDmSqlParserListener) ExitQuery_rewrite(ctx *Query_rewriteContext) {}

// EnterBuild_clause is called when production build_clause is entered.
func (s *BaseDmSqlParserListener) EnterBuild_clause(ctx *Build_clauseContext) {}

// ExitBuild_clause is called when production build_clause is exited.
func (s *BaseDmSqlParserListener) ExitBuild_clause(ctx *Build_clauseContext) {}

// EnterMv_refresh_option is called when production mv_refresh_option is entered.
func (s *BaseDmSqlParserListener) EnterMv_refresh_option(ctx *Mv_refresh_optionContext) {}

// ExitMv_refresh_option is called when production mv_refresh_option is exited.
func (s *BaseDmSqlParserListener) ExitMv_refresh_option(ctx *Mv_refresh_optionContext) {}

// EnterMv_refresh_option_list is called when production mv_refresh_option_list is entered.
func (s *BaseDmSqlParserListener) EnterMv_refresh_option_list(ctx *Mv_refresh_option_listContext) {}

// ExitMv_refresh_option_list is called when production mv_refresh_option_list is exited.
func (s *BaseDmSqlParserListener) ExitMv_refresh_option_list(ctx *Mv_refresh_option_listContext) {}

// EnterMv_refresh_clause is called when production mv_refresh_clause is entered.
func (s *BaseDmSqlParserListener) EnterMv_refresh_clause(ctx *Mv_refresh_clauseContext) {}

// ExitMv_refresh_clause is called when production mv_refresh_clause is exited.
func (s *BaseDmSqlParserListener) ExitMv_refresh_clause(ctx *Mv_refresh_clauseContext) {}

// EnterMv_log_purge_syn_asyn_clause is called when production mv_log_purge_syn_asyn_clause is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_purge_syn_asyn_clause(ctx *Mv_log_purge_syn_asyn_clauseContext) {
}

// ExitMv_log_purge_syn_asyn_clause is called when production mv_log_purge_syn_asyn_clause is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_purge_syn_asyn_clause(ctx *Mv_log_purge_syn_asyn_clauseContext) {
}

// EnterMv_log_purge_clause is called when production mv_log_purge_clause is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) {}

// ExitMv_log_purge_clause is called when production mv_log_purge_clause is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) {}

// EnterMv_log_with_syntax_item is called when production mv_log_with_syntax_item is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_with_syntax_item(ctx *Mv_log_with_syntax_itemContext) {}

// ExitMv_log_with_syntax_item is called when production mv_log_with_syntax_item is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_with_syntax_item(ctx *Mv_log_with_syntax_itemContext) {}

// EnterMv_log_with_syntax_item_list is called when production mv_log_with_syntax_item_list is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_with_syntax_item_list(ctx *Mv_log_with_syntax_item_listContext) {
}

// ExitMv_log_with_syntax_item_list is called when production mv_log_with_syntax_item_list is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_with_syntax_item_list(ctx *Mv_log_with_syntax_item_listContext) {
}

// EnterMv_log_including_new_values is called when production mv_log_including_new_values is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_including_new_values(ctx *Mv_log_including_new_valuesContext) {
}

// ExitMv_log_including_new_values is called when production mv_log_including_new_values is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_including_new_values(ctx *Mv_log_including_new_valuesContext) {
}

// EnterMv_log_with_clause_null is called when production mv_log_with_clause_null is entered.
func (s *BaseDmSqlParserListener) EnterMv_log_with_clause_null(ctx *Mv_log_with_clause_nullContext) {}

// ExitMv_log_with_clause_null is called when production mv_log_with_clause_null is exited.
func (s *BaseDmSqlParserListener) ExitMv_log_with_clause_null(ctx *Mv_log_with_clause_nullContext) {}

// EnterCreate_materialized_view_log_stmt is called when production create_materialized_view_log_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_materialized_view_log_stmt(ctx *Create_materialized_view_log_stmtContext) {
}

// ExitCreate_materialized_view_log_stmt is called when production create_materialized_view_log_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_materialized_view_log_stmt(ctx *Create_materialized_view_log_stmtContext) {
}

// EnterPrebuilt_table_clause_null is called when production prebuilt_table_clause_null is entered.
func (s *BaseDmSqlParserListener) EnterPrebuilt_table_clause_null(ctx *Prebuilt_table_clause_nullContext) {
}

// ExitPrebuilt_table_clause_null is called when production prebuilt_table_clause_null is exited.
func (s *BaseDmSqlParserListener) ExitPrebuilt_table_clause_null(ctx *Prebuilt_table_clause_nullContext) {
}

// EnterCreate_materialized_view_stmt is called when production create_materialized_view_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) {
}

// ExitCreate_materialized_view_stmt is called when production create_materialized_view_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_materialized_view_stmt(ctx *Create_materialized_view_stmtContext) {
}

// EnterCreate_view_stmt is called when production create_view_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_view_stmt(ctx *Create_view_stmtContext) {}

// ExitCreate_view_stmt is called when production create_view_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_view_stmt(ctx *Create_view_stmtContext) {}

// EnterCreate_view_stmt_body is called when production create_view_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterCreate_view_stmt_body(ctx *Create_view_stmt_bodyContext) {}

// ExitCreate_view_stmt_body is called when production create_view_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitCreate_view_stmt_body(ctx *Create_view_stmt_bodyContext) {}

// EnterColumn_list3_options is called when production column_list3_options is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list3_options(ctx *Column_list3_optionsContext) {}

// ExitColumn_list3_options is called when production column_list3_options is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list3_options(ctx *Column_list3_optionsContext) {}

// EnterColumn_list3 is called when production column_list3 is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list3(ctx *Column_list3Context) {}

// ExitColumn_list3 is called when production column_list3 is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list3(ctx *Column_list3Context) {}

// EnterView_column_constraint_def is called when production view_column_constraint_def is entered.
func (s *BaseDmSqlParserListener) EnterView_column_constraint_def(ctx *View_column_constraint_defContext) {
}

// ExitView_column_constraint_def is called when production view_column_constraint_def is exited.
func (s *BaseDmSqlParserListener) ExitView_column_constraint_def(ctx *View_column_constraint_defContext) {
}

// EnterView_column_constraints is called when production view_column_constraints is entered.
func (s *BaseDmSqlParserListener) EnterView_column_constraints(ctx *View_column_constraintsContext) {}

// ExitView_column_constraints is called when production view_column_constraints is exited.
func (s *BaseDmSqlParserListener) ExitView_column_constraints(ctx *View_column_constraintsContext) {}

// EnterView_column_constraint is called when production view_column_constraint is entered.
func (s *BaseDmSqlParserListener) EnterView_column_constraint(ctx *View_column_constraintContext) {}

// ExitView_column_constraint is called when production view_column_constraint is exited.
func (s *BaseDmSqlParserListener) ExitView_column_constraint(ctx *View_column_constraintContext) {}

// EnterView_column_constraint_action is called when production view_column_constraint_action is entered.
func (s *BaseDmSqlParserListener) EnterView_column_constraint_action(ctx *View_column_constraint_actionContext) {
}

// ExitView_column_constraint_action is called when production view_column_constraint_action is exited.
func (s *BaseDmSqlParserListener) ExitView_column_constraint_action(ctx *View_column_constraint_actionContext) {
}

// EnterView_constraint_def is called when production view_constraint_def is entered.
func (s *BaseDmSqlParserListener) EnterView_constraint_def(ctx *View_constraint_defContext) {}

// ExitView_constraint_def is called when production view_constraint_def is exited.
func (s *BaseDmSqlParserListener) ExitView_constraint_def(ctx *View_constraint_defContext) {}

// EnterWith_schemabinding is called when production with_schemabinding is entered.
func (s *BaseDmSqlParserListener) EnterWith_schemabinding(ctx *With_schemabindingContext) {}

// ExitWith_schemabinding is called when production with_schemabinding is exited.
func (s *BaseDmSqlParserListener) ExitWith_schemabinding(ctx *With_schemabindingContext) {}

// EnterColumn_list_options is called when production column_list_options is entered.
func (s *BaseDmSqlParserListener) EnterColumn_list_options(ctx *Column_list_optionsContext) {}

// ExitColumn_list_options is called when production column_list_options is exited.
func (s *BaseDmSqlParserListener) ExitColumn_list_options(ctx *Column_list_optionsContext) {}

// EnterWith_check_or_readonly_option is called when production with_check_or_readonly_option is entered.
func (s *BaseDmSqlParserListener) EnterWith_check_or_readonly_option(ctx *With_check_or_readonly_optionContext) {
}

// ExitWith_check_or_readonly_option is called when production with_check_or_readonly_option is exited.
func (s *BaseDmSqlParserListener) ExitWith_check_or_readonly_option(ctx *With_check_or_readonly_optionContext) {
}

// EnterCheck_level_option is called when production check_level_option is entered.
func (s *BaseDmSqlParserListener) EnterCheck_level_option(ctx *Check_level_optionContext) {}

// ExitCheck_level_option is called when production check_level_option is exited.
func (s *BaseDmSqlParserListener) ExitCheck_level_option(ctx *Check_level_optionContext) {}

// EnterDecl_cursor is called when production decl_cursor is entered.
func (s *BaseDmSqlParserListener) EnterDecl_cursor(ctx *Decl_cursorContext) {}

// ExitDecl_cursor is called when production decl_cursor is exited.
func (s *BaseDmSqlParserListener) ExitDecl_cursor(ctx *Decl_cursorContext) {}

// EnterDelete_stmt is called when production delete_stmt is entered.
func (s *BaseDmSqlParserListener) EnterDelete_stmt(ctx *Delete_stmtContext) {}

// ExitDelete_stmt is called when production delete_stmt is exited.
func (s *BaseDmSqlParserListener) ExitDelete_stmt(ctx *Delete_stmtContext) {}

// EnterDelete_stmt_body is called when production delete_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterDelete_stmt_body(ctx *Delete_stmt_bodyContext) {}

// ExitDelete_stmt_body is called when production delete_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitDelete_stmt_body(ctx *Delete_stmt_bodyContext) {}

// EnterDelete_multi_tv_option is called when production delete_multi_tv_option is entered.
func (s *BaseDmSqlParserListener) EnterDelete_multi_tv_option(ctx *Delete_multi_tv_optionContext) {}

// ExitDelete_multi_tv_option is called when production delete_multi_tv_option is exited.
func (s *BaseDmSqlParserListener) ExitDelete_multi_tv_option(ctx *Delete_multi_tv_optionContext) {}

// EnterCheck_limit_option is called when production check_limit_option is entered.
func (s *BaseDmSqlParserListener) EnterCheck_limit_option(ctx *Check_limit_optionContext) {}

// ExitCheck_limit_option is called when production check_limit_option is exited.
func (s *BaseDmSqlParserListener) ExitCheck_limit_option(ctx *Check_limit_optionContext) {}

// EnterWhere_current_option is called when production where_current_option is entered.
func (s *BaseDmSqlParserListener) EnterWhere_current_option(ctx *Where_current_optionContext) {}

// ExitWhere_current_option is called when production where_current_option is exited.
func (s *BaseDmSqlParserListener) ExitWhere_current_option(ctx *Where_current_optionContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseDmSqlParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseDmSqlParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterStart_with_clause_null is called when production start_with_clause_null is entered.
func (s *BaseDmSqlParserListener) EnterStart_with_clause_null(ctx *Start_with_clause_nullContext) {}

// ExitStart_with_clause_null is called when production start_with_clause_null is exited.
func (s *BaseDmSqlParserListener) ExitStart_with_clause_null(ctx *Start_with_clause_nullContext) {}

// EnterConnect_by_item is called when production connect_by_item is entered.
func (s *BaseDmSqlParserListener) EnterConnect_by_item(ctx *Connect_by_itemContext) {}

// ExitConnect_by_item is called when production connect_by_item is exited.
func (s *BaseDmSqlParserListener) ExitConnect_by_item(ctx *Connect_by_itemContext) {}

// EnterConnect_by_clause is called when production connect_by_clause is entered.
func (s *BaseDmSqlParserListener) EnterConnect_by_clause(ctx *Connect_by_clauseContext) {}

// ExitConnect_by_clause is called when production connect_by_clause is exited.
func (s *BaseDmSqlParserListener) ExitConnect_by_clause(ctx *Connect_by_clauseContext) {}

// EnterHierarchical_query_clause is called when production hierarchical_query_clause is entered.
func (s *BaseDmSqlParserListener) EnterHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) {
}

// ExitHierarchical_query_clause is called when production hierarchical_query_clause is exited.
func (s *BaseDmSqlParserListener) ExitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) {
}

// EnterNocycle_flag is called when production nocycle_flag is entered.
func (s *BaseDmSqlParserListener) EnterNocycle_flag(ctx *Nocycle_flagContext) {}

// ExitNocycle_flag is called when production nocycle_flag is exited.
func (s *BaseDmSqlParserListener) ExitNocycle_flag(ctx *Nocycle_flagContext) {}

// EnterSearch_condition is called when production search_condition is entered.
func (s *BaseDmSqlParserListener) EnterSearch_condition(ctx *Search_conditionContext) {}

// ExitSearch_condition is called when production search_condition is exited.
func (s *BaseDmSqlParserListener) ExitSearch_condition(ctx *Search_conditionContext) {}

// EnterDisconnect_stmt is called when production disconnect_stmt is entered.
func (s *BaseDmSqlParserListener) EnterDisconnect_stmt(ctx *Disconnect_stmtContext) {}

// ExitDisconnect_stmt is called when production disconnect_stmt is exited.
func (s *BaseDmSqlParserListener) ExitDisconnect_stmt(ctx *Disconnect_stmtContext) {}

// EnterDisconnect_option is called when production disconnect_option is entered.
func (s *BaseDmSqlParserListener) EnterDisconnect_option(ctx *Disconnect_optionContext) {}

// ExitDisconnect_option is called when production disconnect_option is exited.
func (s *BaseDmSqlParserListener) ExitDisconnect_option(ctx *Disconnect_optionContext) {}

// EnterDrop_stmt is called when production drop_stmt is entered.
func (s *BaseDmSqlParserListener) EnterDrop_stmt(ctx *Drop_stmtContext) {}

// ExitDrop_stmt is called when production drop_stmt is exited.
func (s *BaseDmSqlParserListener) ExitDrop_stmt(ctx *Drop_stmtContext) {}

// EnterDrop_stmt_body_1 is called when production drop_stmt_body_1 is entered.
func (s *BaseDmSqlParserListener) EnterDrop_stmt_body_1(ctx *Drop_stmt_body_1Context) {}

// ExitDrop_stmt_body_1 is called when production drop_stmt_body_1 is exited.
func (s *BaseDmSqlParserListener) ExitDrop_stmt_body_1(ctx *Drop_stmt_body_1Context) {}

// EnterDrop_stmt_2 is called when production drop_stmt_2 is entered.
func (s *BaseDmSqlParserListener) EnterDrop_stmt_2(ctx *Drop_stmt_2Context) {}

// ExitDrop_stmt_2 is called when production drop_stmt_2 is exited.
func (s *BaseDmSqlParserListener) ExitDrop_stmt_2(ctx *Drop_stmt_2Context) {}

// EnterDrop_id_db_object is called when production drop_id_db_object is entered.
func (s *BaseDmSqlParserListener) EnterDrop_id_db_object(ctx *Drop_id_db_objectContext) {}

// ExitDrop_id_db_object is called when production drop_id_db_object is exited.
func (s *BaseDmSqlParserListener) ExitDrop_id_db_object(ctx *Drop_id_db_objectContext) {}

// EnterId_db_object is called when production id_db_object is entered.
func (s *BaseDmSqlParserListener) EnterId_db_object(ctx *Id_db_objectContext) {}

// ExitId_db_object is called when production id_db_object is exited.
func (s *BaseDmSqlParserListener) ExitId_db_object(ctx *Id_db_objectContext) {}

// EnterDrop_db_object is called when production drop_db_object is entered.
func (s *BaseDmSqlParserListener) EnterDrop_db_object(ctx *Drop_db_objectContext) {}

// ExitDrop_db_object is called when production drop_db_object is exited.
func (s *BaseDmSqlParserListener) ExitDrop_db_object(ctx *Drop_db_objectContext) {}

// EnterExist is called when production exist is entered.
func (s *BaseDmSqlParserListener) EnterExist(ctx *ExistContext) {}

// ExitExist is called when production exist is exited.
func (s *BaseDmSqlParserListener) ExitExist(ctx *ExistContext) {}

// EnterNot_exist is called when production not_exist is entered.
func (s *BaseDmSqlParserListener) EnterNot_exist(ctx *Not_existContext) {}

// ExitNot_exist is called when production not_exist is exited.
func (s *BaseDmSqlParserListener) ExitNot_exist(ctx *Not_existContext) {}

// EnterDb_object is called when production db_object is entered.
func (s *BaseDmSqlParserListener) EnterDb_object(ctx *Db_objectContext) {}

// ExitDb_object is called when production db_object is exited.
func (s *BaseDmSqlParserListener) ExitDb_object(ctx *Db_objectContext) {}

// EnterIs_detach is called when production is_detach is entered.
func (s *BaseDmSqlParserListener) EnterIs_detach(ctx *Is_detachContext) {}

// ExitIs_detach is called when production is_detach is exited.
func (s *BaseDmSqlParserListener) ExitIs_detach(ctx *Is_detachContext) {}

// EnterPurge_option is called when production purge_option is entered.
func (s *BaseDmSqlParserListener) EnterPurge_option(ctx *Purge_optionContext) {}

// ExitPurge_option is called when production purge_option is exited.
func (s *BaseDmSqlParserListener) ExitPurge_option(ctx *Purge_optionContext) {}

// EnterAlter_database_stmt is called when production alter_database_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_database_stmt(ctx *Alter_database_stmtContext) {}

// ExitAlter_database_stmt is called when production alter_database_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_database_stmt(ctx *Alter_database_stmtContext) {}

// EnterAlter_system_action is called when production alter_system_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_system_action(ctx *Alter_system_actionContext) {}

// ExitAlter_system_action is called when production alter_system_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_system_action(ctx *Alter_system_actionContext) {}

// EnterAlter_database_action is called when production alter_database_action is entered.
func (s *BaseDmSqlParserListener) EnterAlter_database_action(ctx *Alter_database_actionContext) {}

// ExitAlter_database_action is called when production alter_database_action is exited.
func (s *BaseDmSqlParserListener) ExitAlter_database_action(ctx *Alter_database_actionContext) {}

// EnterForce is called when production force is entered.
func (s *BaseDmSqlParserListener) EnterForce(ctx *ForceContext) {}

// ExitForce is called when production force is exited.
func (s *BaseDmSqlParserListener) ExitForce(ctx *ForceContext) {}

// EnterTablespace_name is called when production tablespace_name is entered.
func (s *BaseDmSqlParserListener) EnterTablespace_name(ctx *Tablespace_nameContext) {}

// ExitTablespace_name is called when production tablespace_name is exited.
func (s *BaseDmSqlParserListener) ExitTablespace_name(ctx *Tablespace_nameContext) {}

// EnterRaft_name is called when production raft_name is entered.
func (s *BaseDmSqlParserListener) EnterRaft_name(ctx *Raft_nameContext) {}

// ExitRaft_name is called when production raft_name is exited.
func (s *BaseDmSqlParserListener) ExitRaft_name(ctx *Raft_nameContext) {}

// EnterFetch_into is called when production fetch_into is entered.
func (s *BaseDmSqlParserListener) EnterFetch_into(ctx *Fetch_intoContext) {}

// ExitFetch_into is called when production fetch_into is exited.
func (s *BaseDmSqlParserListener) ExitFetch_into(ctx *Fetch_intoContext) {}

// EnterBulk_or_single_into is called when production bulk_or_single_into is entered.
func (s *BaseDmSqlParserListener) EnterBulk_or_single_into(ctx *Bulk_or_single_intoContext) {}

// ExitBulk_or_single_into is called when production bulk_or_single_into is exited.
func (s *BaseDmSqlParserListener) ExitBulk_or_single_into(ctx *Bulk_or_single_intoContext) {}

// EnterFetch_stmt is called when production fetch_stmt is entered.
func (s *BaseDmSqlParserListener) EnterFetch_stmt(ctx *Fetch_stmtContext) {}

// ExitFetch_stmt is called when production fetch_stmt is exited.
func (s *BaseDmSqlParserListener) ExitFetch_stmt(ctx *Fetch_stmtContext) {}

// EnterFetch_statement is called when production fetch_statement is entered.
func (s *BaseDmSqlParserListener) EnterFetch_statement(ctx *Fetch_statementContext) {}

// ExitFetch_statement is called when production fetch_statement is exited.
func (s *BaseDmSqlParserListener) ExitFetch_statement(ctx *Fetch_statementContext) {}

// EnterFetch_tail is called when production fetch_tail is entered.
func (s *BaseDmSqlParserListener) EnterFetch_tail(ctx *Fetch_tailContext) {}

// ExitFetch_tail is called when production fetch_tail is exited.
func (s *BaseDmSqlParserListener) ExitFetch_tail(ctx *Fetch_tailContext) {}

// EnterFetch_limit_option is called when production fetch_limit_option is entered.
func (s *BaseDmSqlParserListener) EnterFetch_limit_option(ctx *Fetch_limit_optionContext) {}

// ExitFetch_limit_option is called when production fetch_limit_option is exited.
func (s *BaseDmSqlParserListener) ExitFetch_limit_option(ctx *Fetch_limit_optionContext) {}

// EnterFetch_option is called when production fetch_option is entered.
func (s *BaseDmSqlParserListener) EnterFetch_option(ctx *Fetch_optionContext) {}

// ExitFetch_option is called when production fetch_option is exited.
func (s *BaseDmSqlParserListener) ExitFetch_option(ctx *Fetch_optionContext) {}

// EnterFetch_direct_option is called when production fetch_direct_option is entered.
func (s *BaseDmSqlParserListener) EnterFetch_direct_option(ctx *Fetch_direct_optionContext) {}

// ExitFetch_direct_option is called when production fetch_direct_option is exited.
func (s *BaseDmSqlParserListener) ExitFetch_direct_option(ctx *Fetch_direct_optionContext) {}

// EnterLog_errors_into is called when production log_errors_into is entered.
func (s *BaseDmSqlParserListener) EnterLog_errors_into(ctx *Log_errors_intoContext) {}

// ExitLog_errors_into is called when production log_errors_into is exited.
func (s *BaseDmSqlParserListener) ExitLog_errors_into(ctx *Log_errors_intoContext) {}

// EnterLog_errors_expression is called when production log_errors_expression is entered.
func (s *BaseDmSqlParserListener) EnterLog_errors_expression(ctx *Log_errors_expressionContext) {}

// ExitLog_errors_expression is called when production log_errors_expression is exited.
func (s *BaseDmSqlParserListener) ExitLog_errors_expression(ctx *Log_errors_expressionContext) {}

// EnterLog_errors_unlimited is called when production log_errors_unlimited is entered.
func (s *BaseDmSqlParserListener) EnterLog_errors_unlimited(ctx *Log_errors_unlimitedContext) {}

// ExitLog_errors_unlimited is called when production log_errors_unlimited is exited.
func (s *BaseDmSqlParserListener) ExitLog_errors_unlimited(ctx *Log_errors_unlimitedContext) {}

// EnterLog_errors is called when production log_errors is entered.
func (s *BaseDmSqlParserListener) EnterLog_errors(ctx *Log_errorsContext) {}

// ExitLog_errors is called when production log_errors is exited.
func (s *BaseDmSqlParserListener) ExitLog_errors(ctx *Log_errorsContext) {}

// EnterInsert_stmt is called when production insert_stmt is entered.
func (s *BaseDmSqlParserListener) EnterInsert_stmt(ctx *Insert_stmtContext) {}

// ExitInsert_stmt is called when production insert_stmt is exited.
func (s *BaseDmSqlParserListener) ExitInsert_stmt(ctx *Insert_stmtContext) {}

// EnterInsert_stmt_body is called when production insert_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterInsert_stmt_body(ctx *Insert_stmt_bodyContext) {}

// ExitInsert_stmt_body is called when production insert_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitInsert_stmt_body(ctx *Insert_stmt_bodyContext) {}

// EnterFull_column_list_options is called when production full_column_list_options is entered.
func (s *BaseDmSqlParserListener) EnterFull_column_list_options(ctx *Full_column_list_optionsContext) {
}

// ExitFull_column_list_options is called when production full_column_list_options is exited.
func (s *BaseDmSqlParserListener) ExitFull_column_list_options(ctx *Full_column_list_optionsContext) {
}

// EnterIns_value_options is called when production ins_value_options is entered.
func (s *BaseDmSqlParserListener) EnterIns_value_options(ctx *Ins_value_optionsContext) {}

// ExitIns_value_options is called when production ins_value_options is exited.
func (s *BaseDmSqlParserListener) ExitIns_value_options(ctx *Ins_value_optionsContext) {}

// EnterInsert_into_single is called when production insert_into_single is entered.
func (s *BaseDmSqlParserListener) EnterInsert_into_single(ctx *Insert_into_singleContext) {}

// ExitInsert_into_single is called when production insert_into_single is exited.
func (s *BaseDmSqlParserListener) ExitInsert_into_single(ctx *Insert_into_singleContext) {}

// EnterMulti_insert_into_list is called when production multi_insert_into_list is entered.
func (s *BaseDmSqlParserListener) EnterMulti_insert_into_list(ctx *Multi_insert_into_listContext) {}

// ExitMulti_insert_into_list is called when production multi_insert_into_list is exited.
func (s *BaseDmSqlParserListener) ExitMulti_insert_into_list(ctx *Multi_insert_into_listContext) {}

// EnterMulti_insert_tag is called when production multi_insert_tag is entered.
func (s *BaseDmSqlParserListener) EnterMulti_insert_tag(ctx *Multi_insert_tagContext) {}

// ExitMulti_insert_tag is called when production multi_insert_tag is exited.
func (s *BaseDmSqlParserListener) ExitMulti_insert_tag(ctx *Multi_insert_tagContext) {}

// EnterInsert_into_single_condition is called when production insert_into_single_condition is entered.
func (s *BaseDmSqlParserListener) EnterInsert_into_single_condition(ctx *Insert_into_single_conditionContext) {
}

// ExitInsert_into_single_condition is called when production insert_into_single_condition is exited.
func (s *BaseDmSqlParserListener) ExitInsert_into_single_condition(ctx *Insert_into_single_conditionContext) {
}

// EnterMulti_insert_into_condition_list is called when production multi_insert_into_condition_list is entered.
func (s *BaseDmSqlParserListener) EnterMulti_insert_into_condition_list(ctx *Multi_insert_into_condition_listContext) {
}

// ExitMulti_insert_into_condition_list is called when production multi_insert_into_condition_list is exited.
func (s *BaseDmSqlParserListener) ExitMulti_insert_into_condition_list(ctx *Multi_insert_into_condition_listContext) {
}

// EnterMulti_insert_into_else is called when production multi_insert_into_else is entered.
func (s *BaseDmSqlParserListener) EnterMulti_insert_into_else(ctx *Multi_insert_into_elseContext) {}

// ExitMulti_insert_into_else is called when production multi_insert_into_else is exited.
func (s *BaseDmSqlParserListener) ExitMulti_insert_into_else(ctx *Multi_insert_into_elseContext) {}

// EnterMulti_insert_stmt_body is called when production multi_insert_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterMulti_insert_stmt_body(ctx *Multi_insert_stmt_bodyContext) {}

// ExitMulti_insert_stmt_body is called when production multi_insert_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitMulti_insert_stmt_body(ctx *Multi_insert_stmt_bodyContext) {}

// EnterInsert_tail is called when production insert_tail is entered.
func (s *BaseDmSqlParserListener) EnterInsert_tail(ctx *Insert_tailContext) {}

// ExitInsert_tail is called when production insert_tail is exited.
func (s *BaseDmSqlParserListener) ExitInsert_tail(ctx *Insert_tailContext) {}

// EnterInsert_action is called when production insert_action is entered.
func (s *BaseDmSqlParserListener) EnterInsert_action(ctx *Insert_actionContext) {}

// ExitInsert_action is called when production insert_action is exited.
func (s *BaseDmSqlParserListener) ExitInsert_action(ctx *Insert_actionContext) {}

// EnterRecord_var_values is called when production record_var_values is entered.
func (s *BaseDmSqlParserListener) EnterRecord_var_values(ctx *Record_var_valuesContext) {}

// ExitRecord_var_values is called when production record_var_values is exited.
func (s *BaseDmSqlParserListener) ExitRecord_var_values(ctx *Record_var_valuesContext) {}

// EnterRecord_var_value is called when production record_var_value is entered.
func (s *BaseDmSqlParserListener) EnterRecord_var_value(ctx *Record_var_valueContext) {}

// ExitRecord_var_value is called when production record_var_value is exited.
func (s *BaseDmSqlParserListener) ExitRecord_var_value(ctx *Record_var_valueContext) {}

// EnterIns_value is called when production ins_value is entered.
func (s *BaseDmSqlParserListener) EnterIns_value(ctx *Ins_valueContext) {}

// ExitIns_value is called when production ins_value is exited.
func (s *BaseDmSqlParserListener) ExitIns_value(ctx *Ins_valueContext) {}

// EnterOpen_stmt is called when production open_stmt is entered.
func (s *BaseDmSqlParserListener) EnterOpen_stmt(ctx *Open_stmtContext) {}

// ExitOpen_stmt is called when production open_stmt is exited.
func (s *BaseDmSqlParserListener) ExitOpen_stmt(ctx *Open_stmtContext) {}

// EnterOpen_statement is called when production open_statement is entered.
func (s *BaseDmSqlParserListener) EnterOpen_statement(ctx *Open_statementContext) {}

// ExitOpen_statement is called when production open_statement is exited.
func (s *BaseDmSqlParserListener) ExitOpen_statement(ctx *Open_statementContext) {}

// EnterOpen_tail is called when production open_tail is entered.
func (s *BaseDmSqlParserListener) EnterOpen_tail(ctx *Open_tailContext) {}

// ExitOpen_tail is called when production open_tail is exited.
func (s *BaseDmSqlParserListener) ExitOpen_tail(ctx *Open_tailContext) {}

// EnterReturn_stmt is called when production return_stmt is entered.
func (s *BaseDmSqlParserListener) EnterReturn_stmt(ctx *Return_stmtContext) {}

// ExitReturn_stmt is called when production return_stmt is exited.
func (s *BaseDmSqlParserListener) ExitReturn_stmt(ctx *Return_stmtContext) {}

// EnterRaise_stmt is called when production raise_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRaise_stmt(ctx *Raise_stmtContext) {}

// ExitRaise_stmt is called when production raise_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRaise_stmt(ctx *Raise_stmtContext) {}

// EnterRollback_stmt is called when production rollback_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRollback_stmt(ctx *Rollback_stmtContext) {}

// ExitRollback_stmt is called when production rollback_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRollback_stmt(ctx *Rollback_stmtContext) {}

// EnterTo_savepoint is called when production to_savepoint is entered.
func (s *BaseDmSqlParserListener) EnterTo_savepoint(ctx *To_savepointContext) {}

// ExitTo_savepoint is called when production to_savepoint is exited.
func (s *BaseDmSqlParserListener) ExitTo_savepoint(ctx *To_savepointContext) {}

// EnterSavepoint_stmt is called when production savepoint_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// ExitSavepoint_stmt is called when production savepoint_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSavepoint_stmt(ctx *Savepoint_stmtContext) {}

// EnterSelect_stmt is called when production select_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSelect_stmt(ctx *Select_stmtContext) {}

// ExitSelect_stmt is called when production select_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSelect_stmt(ctx *Select_stmtContext) {}

// EnterAll_distinct_option is called when production all_distinct_option is entered.
func (s *BaseDmSqlParserListener) EnterAll_distinct_option(ctx *All_distinct_optionContext) {}

// ExitAll_distinct_option is called when production all_distinct_option is exited.
func (s *BaseDmSqlParserListener) ExitAll_distinct_option(ctx *All_distinct_optionContext) {}

// EnterAll_distinct_option_2 is called when production all_distinct_option_2 is entered.
func (s *BaseDmSqlParserListener) EnterAll_distinct_option_2(ctx *All_distinct_option_2Context) {}

// ExitAll_distinct_option_2 is called when production all_distinct_option_2 is exited.
func (s *BaseDmSqlParserListener) ExitAll_distinct_option_2(ctx *All_distinct_option_2Context) {}

// EnterCorresponding_clause is called when production corresponding_clause is entered.
func (s *BaseDmSqlParserListener) EnterCorresponding_clause(ctx *Corresponding_clauseContext) {}

// ExitCorresponding_clause is called when production corresponding_clause is exited.
func (s *BaseDmSqlParserListener) ExitCorresponding_clause(ctx *Corresponding_clauseContext) {}

// EnterTop_option is called when production top_option is entered.
func (s *BaseDmSqlParserListener) EnterTop_option(ctx *Top_optionContext) {}

// ExitTop_option is called when production top_option is exited.
func (s *BaseDmSqlParserListener) ExitTop_option(ctx *Top_optionContext) {}

// EnterLimit_option is called when production limit_option is entered.
func (s *BaseDmSqlParserListener) EnterLimit_option(ctx *Limit_optionContext) {}

// ExitLimit_option is called when production limit_option is exited.
func (s *BaseDmSqlParserListener) ExitLimit_option(ctx *Limit_optionContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseDmSqlParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseDmSqlParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterLimit_not_empty is called when production limit_not_empty is entered.
func (s *BaseDmSqlParserListener) EnterLimit_not_empty(ctx *Limit_not_emptyContext) {}

// ExitLimit_not_empty is called when production limit_not_empty is exited.
func (s *BaseDmSqlParserListener) ExitLimit_not_empty(ctx *Limit_not_emptyContext) {}

// EnterRow_limiting_clause is called when production row_limiting_clause is entered.
func (s *BaseDmSqlParserListener) EnterRow_limiting_clause(ctx *Row_limiting_clauseContext) {}

// ExitRow_limiting_clause is called when production row_limiting_clause is exited.
func (s *BaseDmSqlParserListener) ExitRow_limiting_clause(ctx *Row_limiting_clauseContext) {}

// EnterRow_num_desc is called when production row_num_desc is entered.
func (s *BaseDmSqlParserListener) EnterRow_num_desc(ctx *Row_num_descContext) {}

// ExitRow_num_desc is called when production row_num_desc is exited.
func (s *BaseDmSqlParserListener) ExitRow_num_desc(ctx *Row_num_descContext) {}

// EnterFirst_next_desc is called when production first_next_desc is entered.
func (s *BaseDmSqlParserListener) EnterFirst_next_desc(ctx *First_next_descContext) {}

// ExitFirst_next_desc is called when production first_next_desc is exited.
func (s *BaseDmSqlParserListener) ExitFirst_next_desc(ctx *First_next_descContext) {}

// EnterSelect_item_list is called when production select_item_list is entered.
func (s *BaseDmSqlParserListener) EnterSelect_item_list(ctx *Select_item_listContext) {}

// ExitSelect_item_list is called when production select_item_list is exited.
func (s *BaseDmSqlParserListener) ExitSelect_item_list(ctx *Select_item_listContext) {}

// EnterSelect_item is called when production select_item is entered.
func (s *BaseDmSqlParserListener) EnterSelect_item(ctx *Select_itemContext) {}

// ExitSelect_item is called when production select_item is exited.
func (s *BaseDmSqlParserListener) ExitSelect_item(ctx *Select_itemContext) {}

// EnterAs_alias is called when production as_alias is entered.
func (s *BaseDmSqlParserListener) EnterAs_alias(ctx *As_aliasContext) {}

// ExitAs_alias is called when production as_alias is exited.
func (s *BaseDmSqlParserListener) ExitAs_alias(ctx *As_aliasContext) {}

// EnterSelect_tail is called when production select_tail is entered.
func (s *BaseDmSqlParserListener) EnterSelect_tail(ctx *Select_tailContext) {}

// ExitSelect_tail is called when production select_tail is exited.
func (s *BaseDmSqlParserListener) ExitSelect_tail(ctx *Select_tailContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseDmSqlParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseDmSqlParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_tv_list is called when production from_tv_list is entered.
func (s *BaseDmSqlParserListener) EnterFrom_tv_list(ctx *From_tv_listContext) {}

// ExitFrom_tv_list is called when production from_tv_list is exited.
func (s *BaseDmSqlParserListener) ExitFrom_tv_list(ctx *From_tv_listContext) {}

// EnterFrom_tv is called when production from_tv is entered.
func (s *BaseDmSqlParserListener) EnterFrom_tv(ctx *From_tvContext) {}

// ExitFrom_tv is called when production from_tv is exited.
func (s *BaseDmSqlParserListener) ExitFrom_tv(ctx *From_tvContext) {}

// EnterJoined_table is called when production joined_table is entered.
func (s *BaseDmSqlParserListener) EnterJoined_table(ctx *Joined_tableContext) {}

// ExitJoined_table is called when production joined_table is exited.
func (s *BaseDmSqlParserListener) ExitJoined_table(ctx *Joined_tableContext) {}

// EnterTrxid is called when production trxid is entered.
func (s *BaseDmSqlParserListener) EnterTrxid(ctx *TrxidContext) {}

// ExitTrxid is called when production trxid is exited.
func (s *BaseDmSqlParserListener) ExitTrxid(ctx *TrxidContext) {}

// EnterFlashback_query_low is called when production flashback_query_low is entered.
func (s *BaseDmSqlParserListener) EnterFlashback_query_low(ctx *Flashback_query_lowContext) {}

// ExitFlashback_query_low is called when production flashback_query_low is exited.
func (s *BaseDmSqlParserListener) ExitFlashback_query_low(ctx *Flashback_query_lowContext) {}

// EnterTrxid_option is called when production trxid_option is entered.
func (s *BaseDmSqlParserListener) EnterTrxid_option(ctx *Trxid_optionContext) {}

// ExitTrxid_option is called when production trxid_option is exited.
func (s *BaseDmSqlParserListener) ExitTrxid_option(ctx *Trxid_optionContext) {}

// EnterRange_from_to is called when production range_from_to is entered.
func (s *BaseDmSqlParserListener) EnterRange_from_to(ctx *Range_from_toContext) {}

// ExitRange_from_to is called when production range_from_to is exited.
func (s *BaseDmSqlParserListener) ExitRange_from_to(ctx *Range_from_toContext) {}

// EnterSample_exp is called when production sample_exp is entered.
func (s *BaseDmSqlParserListener) EnterSample_exp(ctx *Sample_expContext) {}

// ExitSample_exp is called when production sample_exp is exited.
func (s *BaseDmSqlParserListener) ExitSample_exp(ctx *Sample_expContext) {}

// EnterPivot_sfun is called when production pivot_sfun is entered.
func (s *BaseDmSqlParserListener) EnterPivot_sfun(ctx *Pivot_sfunContext) {}

// ExitPivot_sfun is called when production pivot_sfun is exited.
func (s *BaseDmSqlParserListener) ExitPivot_sfun(ctx *Pivot_sfunContext) {}

// EnterPivot_sfun_lst is called when production pivot_sfun_lst is entered.
func (s *BaseDmSqlParserListener) EnterPivot_sfun_lst(ctx *Pivot_sfun_lstContext) {}

// ExitPivot_sfun_lst is called when production pivot_sfun_lst is exited.
func (s *BaseDmSqlParserListener) ExitPivot_sfun_lst(ctx *Pivot_sfun_lstContext) {}

// EnterPivot_for_clause is called when production pivot_for_clause is entered.
func (s *BaseDmSqlParserListener) EnterPivot_for_clause(ctx *Pivot_for_clauseContext) {}

// ExitPivot_for_clause is called when production pivot_for_clause is exited.
func (s *BaseDmSqlParserListener) ExitPivot_for_clause(ctx *Pivot_for_clauseContext) {}

// EnterPivot_in_clause1_expr is called when production pivot_in_clause1_expr is entered.
func (s *BaseDmSqlParserListener) EnterPivot_in_clause1_expr(ctx *Pivot_in_clause1_exprContext) {}

// ExitPivot_in_clause1_expr is called when production pivot_in_clause1_expr is exited.
func (s *BaseDmSqlParserListener) ExitPivot_in_clause1_expr(ctx *Pivot_in_clause1_exprContext) {}

// EnterPivot_in_clause_low_1 is called when production pivot_in_clause_low_1 is entered.
func (s *BaseDmSqlParserListener) EnterPivot_in_clause_low_1(ctx *Pivot_in_clause_low_1Context) {}

// ExitPivot_in_clause_low_1 is called when production pivot_in_clause_low_1 is exited.
func (s *BaseDmSqlParserListener) ExitPivot_in_clause_low_1(ctx *Pivot_in_clause_low_1Context) {}

// EnterPivot_in_clause_low_2 is called when production pivot_in_clause_low_2 is entered.
func (s *BaseDmSqlParserListener) EnterPivot_in_clause_low_2(ctx *Pivot_in_clause_low_2Context) {}

// ExitPivot_in_clause_low_2 is called when production pivot_in_clause_low_2 is exited.
func (s *BaseDmSqlParserListener) ExitPivot_in_clause_low_2(ctx *Pivot_in_clause_low_2Context) {}

// EnterPivot_in_clause_low is called when production pivot_in_clause_low is entered.
func (s *BaseDmSqlParserListener) EnterPivot_in_clause_low(ctx *Pivot_in_clause_lowContext) {}

// ExitPivot_in_clause_low is called when production pivot_in_clause_low is exited.
func (s *BaseDmSqlParserListener) ExitPivot_in_clause_low(ctx *Pivot_in_clause_lowContext) {}

// EnterPivot_xml is called when production pivot_xml is entered.
func (s *BaseDmSqlParserListener) EnterPivot_xml(ctx *Pivot_xmlContext) {}

// ExitPivot_xml is called when production pivot_xml is exited.
func (s *BaseDmSqlParserListener) ExitPivot_xml(ctx *Pivot_xmlContext) {}

// EnterPivot_clause_low is called when production pivot_clause_low is entered.
func (s *BaseDmSqlParserListener) EnterPivot_clause_low(ctx *Pivot_clause_lowContext) {}

// ExitPivot_clause_low is called when production pivot_clause_low is exited.
func (s *BaseDmSqlParserListener) ExitPivot_clause_low(ctx *Pivot_clause_lowContext) {}

// EnterUnpivot_val_col_lst is called when production unpivot_val_col_lst is entered.
func (s *BaseDmSqlParserListener) EnterUnpivot_val_col_lst(ctx *Unpivot_val_col_lstContext) {}

// ExitUnpivot_val_col_lst is called when production unpivot_val_col_lst is exited.
func (s *BaseDmSqlParserListener) ExitUnpivot_val_col_lst(ctx *Unpivot_val_col_lstContext) {}

// EnterInclude_clause is called when production include_clause is entered.
func (s *BaseDmSqlParserListener) EnterInclude_clause(ctx *Include_clauseContext) {}

// ExitInclude_clause is called when production include_clause is exited.
func (s *BaseDmSqlParserListener) ExitInclude_clause(ctx *Include_clauseContext) {}

// EnterUnpivot_in_clause_expr is called when production unpivot_in_clause_expr is entered.
func (s *BaseDmSqlParserListener) EnterUnpivot_in_clause_expr(ctx *Unpivot_in_clause_exprContext) {}

// ExitUnpivot_in_clause_expr is called when production unpivot_in_clause_expr is exited.
func (s *BaseDmSqlParserListener) ExitUnpivot_in_clause_expr(ctx *Unpivot_in_clause_exprContext) {}

// EnterUnpivot_in_clause_low is called when production unpivot_in_clause_low is entered.
func (s *BaseDmSqlParserListener) EnterUnpivot_in_clause_low(ctx *Unpivot_in_clause_lowContext) {}

// ExitUnpivot_in_clause_low is called when production unpivot_in_clause_low is exited.
func (s *BaseDmSqlParserListener) ExitUnpivot_in_clause_low(ctx *Unpivot_in_clause_lowContext) {}

// EnterUnpivot_clause_low is called when production unpivot_clause_low is entered.
func (s *BaseDmSqlParserListener) EnterUnpivot_clause_low(ctx *Unpivot_clause_lowContext) {}

// ExitUnpivot_clause_low is called when production unpivot_clause_low is exited.
func (s *BaseDmSqlParserListener) ExitUnpivot_clause_low(ctx *Unpivot_clause_lowContext) {}

// EnterSample_clause_low is called when production sample_clause_low is entered.
func (s *BaseDmSqlParserListener) EnterSample_clause_low(ctx *Sample_clause_lowContext) {}

// ExitSample_clause_low is called when production sample_clause_low is exited.
func (s *BaseDmSqlParserListener) ExitSample_clause_low(ctx *Sample_clause_lowContext) {}

// EnterNormal_tv_name is called when production normal_tv_name is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_name(ctx *Normal_tv_nameContext) {}

// ExitNormal_tv_name is called when production normal_tv_name is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_name(ctx *Normal_tv_nameContext) {}

// EnterNormal_tv_tail is called when production normal_tv_tail is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_tail(ctx *Normal_tv_tailContext) {}

// ExitNormal_tv_tail is called when production normal_tv_tail is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_tail(ctx *Normal_tv_tailContext) {}

// EnterNormal_tv_tail_low is called when production normal_tv_tail_low is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_tail_low(ctx *Normal_tv_tail_lowContext) {}

// ExitNormal_tv_tail_low is called when production normal_tv_tail_low is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_tail_low(ctx *Normal_tv_tail_lowContext) {}

// EnterNormal_alias is called when production normal_alias is entered.
func (s *BaseDmSqlParserListener) EnterNormal_alias(ctx *Normal_aliasContext) {}

// ExitNormal_alias is called when production normal_alias is exited.
func (s *BaseDmSqlParserListener) ExitNormal_alias(ctx *Normal_aliasContext) {}

// EnterNormal_tv_tail_low2 is called when production normal_tv_tail_low2 is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_tail_low2(ctx *Normal_tv_tail_low2Context) {}

// ExitNormal_tv_tail_low2 is called when production normal_tv_tail_low2 is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_tail_low2(ctx *Normal_tv_tail_low2Context) {}

// EnterNormal_tv_tail_low3 is called when production normal_tv_tail_low3 is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_tail_low3(ctx *Normal_tv_tail_low3Context) {}

// ExitNormal_tv_tail_low3 is called when production normal_tv_tail_low3 is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_tail_low3(ctx *Normal_tv_tail_low3Context) {}

// EnterNormal_tv_derived_table_options is called when production normal_tv_derived_table_options is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_derived_table_options(ctx *Normal_tv_derived_table_optionsContext) {
}

// ExitNormal_tv_derived_table_options is called when production normal_tv_derived_table_options is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_derived_table_options(ctx *Normal_tv_derived_table_optionsContext) {
}

// EnterNormal_tv_derived_table_low is called when production normal_tv_derived_table_low is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_derived_table_low(ctx *Normal_tv_derived_table_lowContext) {
}

// ExitNormal_tv_derived_table_low is called when production normal_tv_derived_table_low is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_derived_table_low(ctx *Normal_tv_derived_table_lowContext) {
}

// EnterNormal_tv_derived_table is called when production normal_tv_derived_table is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv_derived_table(ctx *Normal_tv_derived_tableContext) {}

// ExitNormal_tv_derived_table is called when production normal_tv_derived_table is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv_derived_table(ctx *Normal_tv_derived_tableContext) {}

// EnterSelect_with_paran_with_alias is called when production select_with_paran_with_alias is entered.
func (s *BaseDmSqlParserListener) EnterSelect_with_paran_with_alias(ctx *Select_with_paran_with_aliasContext) {
}

// ExitSelect_with_paran_with_alias is called when production select_with_paran_with_alias is exited.
func (s *BaseDmSqlParserListener) ExitSelect_with_paran_with_alias(ctx *Select_with_paran_with_aliasContext) {
}

// EnterFrom_table_exp is called when production from_table_exp is entered.
func (s *BaseDmSqlParserListener) EnterFrom_table_exp(ctx *From_table_expContext) {}

// ExitFrom_table_exp is called when production from_table_exp is exited.
func (s *BaseDmSqlParserListener) ExitFrom_table_exp(ctx *From_table_expContext) {}

// EnterFrom_table_select_with_paran is called when production from_table_select_with_paran is entered.
func (s *BaseDmSqlParserListener) EnterFrom_table_select_with_paran(ctx *From_table_select_with_paranContext) {
}

// ExitFrom_table_select_with_paran is called when production from_table_select_with_paran is exited.
func (s *BaseDmSqlParserListener) ExitFrom_table_select_with_paran(ctx *From_table_select_with_paranContext) {
}

// EnterNormal_tv is called when production normal_tv is entered.
func (s *BaseDmSqlParserListener) EnterNormal_tv(ctx *Normal_tvContext) {}

// ExitNormal_tv is called when production normal_tv is exited.
func (s *BaseDmSqlParserListener) ExitNormal_tv(ctx *Normal_tvContext) {}

// EnterXml_passing is called when production xml_passing is entered.
func (s *BaseDmSqlParserListener) EnterXml_passing(ctx *Xml_passingContext) {}

// ExitXml_passing is called when production xml_passing is exited.
func (s *BaseDmSqlParserListener) ExitXml_passing(ctx *Xml_passingContext) {}

// EnterXmlcoldef_lst_options is called when production xmlcoldef_lst_options is entered.
func (s *BaseDmSqlParserListener) EnterXmlcoldef_lst_options(ctx *Xmlcoldef_lst_optionsContext) {}

// ExitXmlcoldef_lst_options is called when production xmlcoldef_lst_options is exited.
func (s *BaseDmSqlParserListener) ExitXmlcoldef_lst_options(ctx *Xmlcoldef_lst_optionsContext) {}

// EnterXmlcoldef_lst is called when production xmlcoldef_lst is entered.
func (s *BaseDmSqlParserListener) EnterXmlcoldef_lst(ctx *Xmlcoldef_lstContext) {}

// ExitXmlcoldef_lst is called when production xmlcoldef_lst is exited.
func (s *BaseDmSqlParserListener) ExitXmlcoldef_lst(ctx *Xmlcoldef_lstContext) {}

// EnterXmlcoldef is called when production xmlcoldef is entered.
func (s *BaseDmSqlParserListener) EnterXmlcoldef(ctx *XmlcoldefContext) {}

// ExitXmlcoldef is called when production xmlcoldef is exited.
func (s *BaseDmSqlParserListener) ExitXmlcoldef(ctx *XmlcoldefContext) {}

// EnterOn_error is called when production on_error is entered.
func (s *BaseDmSqlParserListener) EnterOn_error(ctx *On_errorContext) {}

// ExitOn_error is called when production on_error is exited.
func (s *BaseDmSqlParserListener) ExitOn_error(ctx *On_errorContext) {}

// EnterJsoncol_lst is called when production jsoncol_lst is entered.
func (s *BaseDmSqlParserListener) EnterJsoncol_lst(ctx *Jsoncol_lstContext) {}

// ExitJsoncol_lst is called when production jsoncol_lst is exited.
func (s *BaseDmSqlParserListener) ExitJsoncol_lst(ctx *Jsoncol_lstContext) {}

// EnterJsoncoldef_lst is called when production jsoncoldef_lst is entered.
func (s *BaseDmSqlParserListener) EnterJsoncoldef_lst(ctx *Jsoncoldef_lstContext) {}

// ExitJsoncoldef_lst is called when production jsoncoldef_lst is exited.
func (s *BaseDmSqlParserListener) ExitJsoncoldef_lst(ctx *Jsoncoldef_lstContext) {}

// EnterJsoncoldef is called when production jsoncoldef is entered.
func (s *BaseDmSqlParserListener) EnterJsoncoldef(ctx *JsoncoldefContext) {}

// ExitJsoncoldef is called when production jsoncoldef is exited.
func (s *BaseDmSqlParserListener) ExitJsoncoldef(ctx *JsoncoldefContext) {}

// EnterJson_exists_col is called when production json_exists_col is entered.
func (s *BaseDmSqlParserListener) EnterJson_exists_col(ctx *Json_exists_colContext) {}

// ExitJson_exists_col is called when production json_exists_col is exited.
func (s *BaseDmSqlParserListener) ExitJson_exists_col(ctx *Json_exists_colContext) {}

// EnterJson_qurey_col is called when production json_qurey_col is entered.
func (s *BaseDmSqlParserListener) EnterJson_qurey_col(ctx *Json_qurey_colContext) {}

// ExitJson_qurey_col is called when production json_qurey_col is exited.
func (s *BaseDmSqlParserListener) ExitJson_qurey_col(ctx *Json_qurey_colContext) {}

// EnterJson_value_col is called when production json_value_col is entered.
func (s *BaseDmSqlParserListener) EnterJson_value_col(ctx *Json_value_colContext) {}

// ExitJson_value_col is called when production json_value_col is exited.
func (s *BaseDmSqlParserListener) ExitJson_value_col(ctx *Json_value_colContext) {}

// EnterJson_nested_col is called when production json_nested_col is entered.
func (s *BaseDmSqlParserListener) EnterJson_nested_col(ctx *Json_nested_colContext) {}

// ExitJson_nested_col is called when production json_nested_col is exited.
func (s *BaseDmSqlParserListener) ExitJson_nested_col(ctx *Json_nested_colContext) {}

// EnterOrdinality_col is called when production ordinality_col is entered.
func (s *BaseDmSqlParserListener) EnterOrdinality_col(ctx *Ordinality_colContext) {}

// ExitOrdinality_col is called when production ordinality_col is exited.
func (s *BaseDmSqlParserListener) ExitOrdinality_col(ctx *Ordinality_colContext) {}

// EnterJoined_table_element is called when production joined_table_element is entered.
func (s *BaseDmSqlParserListener) EnterJoined_table_element(ctx *Joined_table_elementContext) {}

// ExitJoined_table_element is called when production joined_table_element is exited.
func (s *BaseDmSqlParserListener) ExitJoined_table_element(ctx *Joined_table_elementContext) {}

// EnterCross_outer_apply_clause is called when production cross_outer_apply_clause is entered.
func (s *BaseDmSqlParserListener) EnterCross_outer_apply_clause(ctx *Cross_outer_apply_clauseContext) {
}

// ExitCross_outer_apply_clause is called when production cross_outer_apply_clause is exited.
func (s *BaseDmSqlParserListener) ExitCross_outer_apply_clause(ctx *Cross_outer_apply_clauseContext) {
}

// EnterCross_outer_apply_join is called when production cross_outer_apply_join is entered.
func (s *BaseDmSqlParserListener) EnterCross_outer_apply_join(ctx *Cross_outer_apply_joinContext) {}

// ExitCross_outer_apply_join is called when production cross_outer_apply_join is exited.
func (s *BaseDmSqlParserListener) ExitCross_outer_apply_join(ctx *Cross_outer_apply_joinContext) {}

// EnterCross_join is called when production cross_join is entered.
func (s *BaseDmSqlParserListener) EnterCross_join(ctx *Cross_joinContext) {}

// ExitCross_join is called when production cross_join is exited.
func (s *BaseDmSqlParserListener) ExitCross_join(ctx *Cross_joinContext) {}

// EnterPartition_out_join is called when production partition_out_join is entered.
func (s *BaseDmSqlParserListener) EnterPartition_out_join(ctx *Partition_out_joinContext) {}

// ExitPartition_out_join is called when production partition_out_join is exited.
func (s *BaseDmSqlParserListener) ExitPartition_out_join(ctx *Partition_out_joinContext) {}

// EnterQualified_join is called when production qualified_join is entered.
func (s *BaseDmSqlParserListener) EnterQualified_join(ctx *Qualified_joinContext) {}

// ExitQualified_join is called when production qualified_join is exited.
func (s *BaseDmSqlParserListener) ExitQualified_join(ctx *Qualified_joinContext) {}

// EnterQualified_joinspec is called when production qualified_joinspec is entered.
func (s *BaseDmSqlParserListener) EnterQualified_joinspec(ctx *Qualified_joinspecContext) {}

// ExitQualified_joinspec is called when production qualified_joinspec is exited.
func (s *BaseDmSqlParserListener) ExitQualified_joinspec(ctx *Qualified_joinspecContext) {}

// EnterNamed_columns_join is called when production named_columns_join is entered.
func (s *BaseDmSqlParserListener) EnterNamed_columns_join(ctx *Named_columns_joinContext) {}

// ExitNamed_columns_join is called when production named_columns_join is exited.
func (s *BaseDmSqlParserListener) ExitNamed_columns_join(ctx *Named_columns_joinContext) {}

// EnterJoin_hint is called when production join_hint is entered.
func (s *BaseDmSqlParserListener) EnterJoin_hint(ctx *Join_hintContext) {}

// ExitJoin_hint is called when production join_hint is exited.
func (s *BaseDmSqlParserListener) ExitJoin_hint(ctx *Join_hintContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BaseDmSqlParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BaseDmSqlParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterOuter_join_type is called when production outer_join_type is entered.
func (s *BaseDmSqlParserListener) EnterOuter_join_type(ctx *Outer_join_typeContext) {}

// ExitOuter_join_type is called when production outer_join_type is exited.
func (s *BaseDmSqlParserListener) ExitOuter_join_type(ctx *Outer_join_typeContext) {}

// EnterJoin_condition is called when production join_condition is entered.
func (s *BaseDmSqlParserListener) EnterJoin_condition(ctx *Join_conditionContext) {}

// ExitJoin_condition is called when production join_condition is exited.
func (s *BaseDmSqlParserListener) ExitJoin_condition(ctx *Join_conditionContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BaseDmSqlParserListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BaseDmSqlParserListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterGroup_item is called when production group_item is entered.
func (s *BaseDmSqlParserListener) EnterGroup_item(ctx *Group_itemContext) {}

// ExitGroup_item is called when production group_item is exited.
func (s *BaseDmSqlParserListener) ExitGroup_item(ctx *Group_itemContext) {}

// EnterExp_rollup_cube_item2 is called when production exp_rollup_cube_item2 is entered.
func (s *BaseDmSqlParserListener) EnterExp_rollup_cube_item2(ctx *Exp_rollup_cube_item2Context) {}

// ExitExp_rollup_cube_item2 is called when production exp_rollup_cube_item2 is exited.
func (s *BaseDmSqlParserListener) ExitExp_rollup_cube_item2(ctx *Exp_rollup_cube_item2Context) {}

// EnterExp_rollup_cube_item is called when production exp_rollup_cube_item is entered.
func (s *BaseDmSqlParserListener) EnterExp_rollup_cube_item(ctx *Exp_rollup_cube_itemContext) {}

// ExitExp_rollup_cube_item is called when production exp_rollup_cube_item is exited.
func (s *BaseDmSqlParserListener) ExitExp_rollup_cube_item(ctx *Exp_rollup_cube_itemContext) {}

// EnterGrouping_set_items is called when production grouping_set_items is entered.
func (s *BaseDmSqlParserListener) EnterGrouping_set_items(ctx *Grouping_set_itemsContext) {}

// ExitGrouping_set_items is called when production grouping_set_items is exited.
func (s *BaseDmSqlParserListener) ExitGrouping_set_items(ctx *Grouping_set_itemsContext) {}

// EnterGrouping_set_item is called when production grouping_set_item is entered.
func (s *BaseDmSqlParserListener) EnterGrouping_set_item(ctx *Grouping_set_itemContext) {}

// ExitGrouping_set_item is called when production grouping_set_item is exited.
func (s *BaseDmSqlParserListener) ExitGrouping_set_item(ctx *Grouping_set_itemContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseDmSqlParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseDmSqlParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterWithout_into_select is called when production without_into_select is entered.
func (s *BaseDmSqlParserListener) EnterWithout_into_select(ctx *Without_into_selectContext) {}

// ExitWithout_into_select is called when production without_into_select is exited.
func (s *BaseDmSqlParserListener) ExitWithout_into_select(ctx *Without_into_selectContext) {}

// EnterSel_clause_app is called when production sel_clause_app is entered.
func (s *BaseDmSqlParserListener) EnterSel_clause_app(ctx *Sel_clause_appContext) {}

// ExitSel_clause_app is called when production sel_clause_app is exited.
func (s *BaseDmSqlParserListener) ExitSel_clause_app(ctx *Sel_clause_appContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseDmSqlParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseDmSqlParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSimple_select is called when production simple_select is entered.
func (s *BaseDmSqlParserListener) EnterSimple_select(ctx *Simple_selectContext) {}

// ExitSimple_select is called when production simple_select is exited.
func (s *BaseDmSqlParserListener) ExitSimple_select(ctx *Simple_selectContext) {}

// EnterSelect_with_paran is called when production select_with_paran is entered.
func (s *BaseDmSqlParserListener) EnterSelect_with_paran(ctx *Select_with_paranContext) {}

// ExitSelect_with_paran is called when production select_with_paran is exited.
func (s *BaseDmSqlParserListener) ExitSelect_with_paran(ctx *Select_with_paranContext) {}

// EnterQuery_exp is called when production query_exp is entered.
func (s *BaseDmSqlParserListener) EnterQuery_exp(ctx *Query_expContext) {}

// ExitQuery_exp is called when production query_exp is exited.
func (s *BaseDmSqlParserListener) ExitQuery_exp(ctx *Query_expContext) {}

// EnterFor_xml_path is called when production for_xml_path is entered.
func (s *BaseDmSqlParserListener) EnterFor_xml_path(ctx *For_xml_pathContext) {}

// ExitFor_xml_path is called when production for_xml_path is exited.
func (s *BaseDmSqlParserListener) ExitFor_xml_path(ctx *For_xml_pathContext) {}

// EnterWith_tag is called when production with_tag is entered.
func (s *BaseDmSqlParserListener) EnterWith_tag(ctx *With_tagContext) {}

// ExitWith_tag is called when production with_tag is exited.
func (s *BaseDmSqlParserListener) ExitWith_tag(ctx *With_tagContext) {}

// EnterWith_option is called when production with_option is entered.
func (s *BaseDmSqlParserListener) EnterWith_option(ctx *With_optionContext) {}

// ExitWith_option is called when production with_option is exited.
func (s *BaseDmSqlParserListener) ExitWith_option(ctx *With_optionContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseDmSqlParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseDmSqlParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterWith_function_list is called when production with_function_list is entered.
func (s *BaseDmSqlParserListener) EnterWith_function_list(ctx *With_function_listContext) {}

// ExitWith_function_list is called when production with_function_list is exited.
func (s *BaseDmSqlParserListener) ExitWith_function_list(ctx *With_function_listContext) {}

// EnterFunc_def_inner is called when production func_def_inner is entered.
func (s *BaseDmSqlParserListener) EnterFunc_def_inner(ctx *Func_def_innerContext) {}

// ExitFunc_def_inner is called when production func_def_inner is exited.
func (s *BaseDmSqlParserListener) ExitFunc_def_inner(ctx *Func_def_innerContext) {}

// EnterWith_function_list_element is called when production with_function_list_element is entered.
func (s *BaseDmSqlParserListener) EnterWith_function_list_element(ctx *With_function_list_elementContext) {
}

// ExitWith_function_list_element is called when production with_function_list_element is exited.
func (s *BaseDmSqlParserListener) ExitWith_function_list_element(ctx *With_function_list_elementContext) {
}

// EnterWith_view_list is called when production with_view_list is entered.
func (s *BaseDmSqlParserListener) EnterWith_view_list(ctx *With_view_listContext) {}

// ExitWith_view_list is called when production with_view_list is exited.
func (s *BaseDmSqlParserListener) ExitWith_view_list(ctx *With_view_listContext) {}

// EnterDepth_type_option is called when production depth_type_option is entered.
func (s *BaseDmSqlParserListener) EnterDepth_type_option(ctx *Depth_type_optionContext) {}

// ExitDepth_type_option is called when production depth_type_option is exited.
func (s *BaseDmSqlParserListener) ExitDepth_type_option(ctx *Depth_type_optionContext) {}

// EnterSearch_clause is called when production search_clause is entered.
func (s *BaseDmSqlParserListener) EnterSearch_clause(ctx *Search_clauseContext) {}

// ExitSearch_clause is called when production search_clause is exited.
func (s *BaseDmSqlParserListener) ExitSearch_clause(ctx *Search_clauseContext) {}

// EnterCycle_clause is called when production cycle_clause is entered.
func (s *BaseDmSqlParserListener) EnterCycle_clause(ctx *Cycle_clauseContext) {}

// ExitCycle_clause is called when production cycle_clause is exited.
func (s *BaseDmSqlParserListener) ExitCycle_clause(ctx *Cycle_clauseContext) {}

// EnterWith_view_list_element is called when production with_view_list_element is entered.
func (s *BaseDmSqlParserListener) EnterWith_view_list_element(ctx *With_view_list_elementContext) {}

// ExitWith_view_list_element is called when production with_view_list_element is exited.
func (s *BaseDmSqlParserListener) ExitWith_view_list_element(ctx *With_view_list_elementContext) {}

// EnterAssignment_obj_list is called when production assignment_obj_list is entered.
func (s *BaseDmSqlParserListener) EnterAssignment_obj_list(ctx *Assignment_obj_listContext) {}

// ExitAssignment_obj_list is called when production assignment_obj_list is exited.
func (s *BaseDmSqlParserListener) ExitAssignment_obj_list(ctx *Assignment_obj_listContext) {}

// EnterAssignment_obj is called when production assignment_obj is entered.
func (s *BaseDmSqlParserListener) EnterAssignment_obj(ctx *Assignment_objContext) {}

// ExitAssignment_obj is called when production assignment_obj is exited.
func (s *BaseDmSqlParserListener) ExitAssignment_obj(ctx *Assignment_objContext) {}

// EnterOrder_by_options is called when production order_by_options is entered.
func (s *BaseDmSqlParserListener) EnterOrder_by_options(ctx *Order_by_optionsContext) {}

// ExitOrder_by_options is called when production order_by_options is exited.
func (s *BaseDmSqlParserListener) ExitOrder_by_options(ctx *Order_by_optionsContext) {}

// EnterOrder_by is called when production order_by is entered.
func (s *BaseDmSqlParserListener) EnterOrder_by(ctx *Order_byContext) {}

// ExitOrder_by is called when production order_by is exited.
func (s *BaseDmSqlParserListener) ExitOrder_by(ctx *Order_byContext) {}

// EnterAsc_desc_option is called when production asc_desc_option is entered.
func (s *BaseDmSqlParserListener) EnterAsc_desc_option(ctx *Asc_desc_optionContext) {}

// ExitAsc_desc_option is called when production asc_desc_option is exited.
func (s *BaseDmSqlParserListener) ExitAsc_desc_option(ctx *Asc_desc_optionContext) {}

// EnterNulls_last_option is called when production nulls_last_option is entered.
func (s *BaseDmSqlParserListener) EnterNulls_last_option(ctx *Nulls_last_optionContext) {}

// ExitNulls_last_option is called when production nulls_last_option is exited.
func (s *BaseDmSqlParserListener) ExitNulls_last_option(ctx *Nulls_last_optionContext) {}

// EnterCollate_option is called when production collate_option is entered.
func (s *BaseDmSqlParserListener) EnterCollate_option(ctx *Collate_optionContext) {}

// ExitCollate_option is called when production collate_option is exited.
func (s *BaseDmSqlParserListener) ExitCollate_option(ctx *Collate_optionContext) {}

// EnterOrder_by_list is called when production order_by_list is entered.
func (s *BaseDmSqlParserListener) EnterOrder_by_list(ctx *Order_by_listContext) {}

// ExitOrder_by_list is called when production order_by_list is exited.
func (s *BaseDmSqlParserListener) ExitOrder_by_list(ctx *Order_by_listContext) {}

// EnterOrder_by_item is called when production order_by_item is entered.
func (s *BaseDmSqlParserListener) EnterOrder_by_item(ctx *Order_by_itemContext) {}

// ExitOrder_by_item is called when production order_by_item is exited.
func (s *BaseDmSqlParserListener) ExitOrder_by_item(ctx *Order_by_itemContext) {}

// EnterFor_update_options is called when production for_update_options is entered.
func (s *BaseDmSqlParserListener) EnterFor_update_options(ctx *For_update_optionsContext) {}

// ExitFor_update_options is called when production for_update_options is exited.
func (s *BaseDmSqlParserListener) ExitFor_update_options(ctx *For_update_optionsContext) {}

// EnterFor_update is called when production for_update is entered.
func (s *BaseDmSqlParserListener) EnterFor_update(ctx *For_updateContext) {}

// ExitFor_update is called when production for_update is exited.
func (s *BaseDmSqlParserListener) ExitFor_update(ctx *For_updateContext) {}

// EnterSet_session_stmt is called when production set_session_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSet_session_stmt(ctx *Set_session_stmtContext) {}

// ExitSet_session_stmt is called when production set_session_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSet_session_stmt(ctx *Set_session_stmtContext) {}

// EnterSet_trans_stmt is called when production set_trans_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSet_trans_stmt(ctx *Set_trans_stmtContext) {}

// ExitSet_trans_stmt is called when production set_trans_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSet_trans_stmt(ctx *Set_trans_stmtContext) {}

// EnterTrans_mode_lstl is called when production trans_mode_lstl is entered.
func (s *BaseDmSqlParserListener) EnterTrans_mode_lstl(ctx *Trans_mode_lstlContext) {}

// ExitTrans_mode_lstl is called when production trans_mode_lstl is exited.
func (s *BaseDmSqlParserListener) ExitTrans_mode_lstl(ctx *Trans_mode_lstlContext) {}

// EnterTrans_mode_lst is called when production trans_mode_lst is entered.
func (s *BaseDmSqlParserListener) EnterTrans_mode_lst(ctx *Trans_mode_lstContext) {}

// ExitTrans_mode_lst is called when production trans_mode_lst is exited.
func (s *BaseDmSqlParserListener) ExitTrans_mode_lst(ctx *Trans_mode_lstContext) {}

// EnterTrans_mode is called when production trans_mode is entered.
func (s *BaseDmSqlParserListener) EnterTrans_mode(ctx *Trans_modeContext) {}

// ExitTrans_mode is called when production trans_mode is exited.
func (s *BaseDmSqlParserListener) ExitTrans_mode(ctx *Trans_modeContext) {}

// EnterTime_zone_exp_new is called when production time_zone_exp_new is entered.
func (s *BaseDmSqlParserListener) EnterTime_zone_exp_new(ctx *Time_zone_exp_newContext) {}

// ExitTime_zone_exp_new is called when production time_zone_exp_new is exited.
func (s *BaseDmSqlParserListener) ExitTime_zone_exp_new(ctx *Time_zone_exp_newContext) {}

// EnterTrans_rw_option is called when production trans_rw_option is entered.
func (s *BaseDmSqlParserListener) EnterTrans_rw_option(ctx *Trans_rw_optionContext) {}

// ExitTrans_rw_option is called when production trans_rw_option is exited.
func (s *BaseDmSqlParserListener) ExitTrans_rw_option(ctx *Trans_rw_optionContext) {}

// EnterTrans_level_option is called when production trans_level_option is entered.
func (s *BaseDmSqlParserListener) EnterTrans_level_option(ctx *Trans_level_optionContext) {}

// ExitTrans_level_option is called when production trans_level_option is exited.
func (s *BaseDmSqlParserListener) ExitTrans_level_option(ctx *Trans_level_optionContext) {}

// EnterLock_table_stmt is called when production lock_table_stmt is entered.
func (s *BaseDmSqlParserListener) EnterLock_table_stmt(ctx *Lock_table_stmtContext) {}

// ExitLock_table_stmt is called when production lock_table_stmt is exited.
func (s *BaseDmSqlParserListener) ExitLock_table_stmt(ctx *Lock_table_stmtContext) {}

// EnterLock_mode_option is called when production lock_mode_option is entered.
func (s *BaseDmSqlParserListener) EnterLock_mode_option(ctx *Lock_mode_optionContext) {}

// ExitLock_mode_option is called when production lock_mode_option is exited.
func (s *BaseDmSqlParserListener) ExitLock_mode_option(ctx *Lock_mode_optionContext) {}

// EnterSet_identins_stmt is called when production set_identins_stmt is entered.
func (s *BaseDmSqlParserListener) EnterSet_identins_stmt(ctx *Set_identins_stmtContext) {}

// ExitSet_identins_stmt is called when production set_identins_stmt is exited.
func (s *BaseDmSqlParserListener) ExitSet_identins_stmt(ctx *Set_identins_stmtContext) {}

// EnterSet_identins_option is called when production set_identins_option is entered.
func (s *BaseDmSqlParserListener) EnterSet_identins_option(ctx *Set_identins_optionContext) {}

// ExitSet_identins_option is called when production set_identins_option is exited.
func (s *BaseDmSqlParserListener) ExitSet_identins_option(ctx *Set_identins_optionContext) {}

// EnterTrunc_table_stmt is called when production trunc_table_stmt is entered.
func (s *BaseDmSqlParserListener) EnterTrunc_table_stmt(ctx *Trunc_table_stmtContext) {}

// ExitTrunc_table_stmt is called when production trunc_table_stmt is exited.
func (s *BaseDmSqlParserListener) ExitTrunc_table_stmt(ctx *Trunc_table_stmtContext) {}

// EnterUpdate_stmt is called when production update_stmt is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_stmt(ctx *Update_stmtContext) {}

// ExitUpdate_stmt is called when production update_stmt is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_stmt(ctx *Update_stmtContext) {}

// EnterUpdate_stmt_body is called when production update_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_stmt_body(ctx *Update_stmt_bodyContext) {}

// ExitUpdate_stmt_body is called when production update_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_stmt_body(ctx *Update_stmt_bodyContext) {}

// EnterUpdate_tv_list is called when production update_tv_list is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_tv_list(ctx *Update_tv_listContext) {}

// ExitUpdate_tv_list is called when production update_tv_list is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_tv_list(ctx *Update_tv_listContext) {}

// EnterReturn_item is called when production return_item is entered.
func (s *BaseDmSqlParserListener) EnterReturn_item(ctx *Return_itemContext) {}

// ExitReturn_item is called when production return_item is exited.
func (s *BaseDmSqlParserListener) ExitReturn_item(ctx *Return_itemContext) {}

// EnterReturn_item_list is called when production return_item_list is entered.
func (s *BaseDmSqlParserListener) EnterReturn_item_list(ctx *Return_item_listContext) {}

// ExitReturn_item_list is called when production return_item_list is exited.
func (s *BaseDmSqlParserListener) ExitReturn_item_list(ctx *Return_item_listContext) {}

// EnterReturn_option is called when production return_option is entered.
func (s *BaseDmSqlParserListener) EnterReturn_option(ctx *Return_optionContext) {}

// ExitReturn_option is called when production return_option is exited.
func (s *BaseDmSqlParserListener) ExitReturn_option(ctx *Return_optionContext) {}

// EnterReturn_into_obj is called when production return_into_obj is entered.
func (s *BaseDmSqlParserListener) EnterReturn_into_obj(ctx *Return_into_objContext) {}

// ExitReturn_into_obj is called when production return_into_obj is exited.
func (s *BaseDmSqlParserListener) ExitReturn_into_obj(ctx *Return_into_objContext) {}

// EnterCollect_into_rset is called when production collect_into_rset is entered.
func (s *BaseDmSqlParserListener) EnterCollect_into_rset(ctx *Collect_into_rsetContext) {}

// ExitCollect_into_rset is called when production collect_into_rset is exited.
func (s *BaseDmSqlParserListener) ExitCollect_into_rset(ctx *Collect_into_rsetContext) {}

// EnterAlias_option is called when production alias_option is entered.
func (s *BaseDmSqlParserListener) EnterAlias_option(ctx *Alias_optionContext) {}

// ExitAlias_option is called when production alias_option is exited.
func (s *BaseDmSqlParserListener) ExitAlias_option(ctx *Alias_optionContext) {}

// EnterSet_value_list is called when production set_value_list is entered.
func (s *BaseDmSqlParserListener) EnterSet_value_list(ctx *Set_value_listContext) {}

// ExitSet_value_list is called when production set_value_list is exited.
func (s *BaseDmSqlParserListener) ExitSet_value_list(ctx *Set_value_listContext) {}

// EnterSet_value_node is called when production set_value_node is entered.
func (s *BaseDmSqlParserListener) EnterSet_value_node(ctx *Set_value_nodeContext) {}

// ExitSet_value_node is called when production set_value_node is exited.
func (s *BaseDmSqlParserListener) ExitSet_value_node(ctx *Set_value_nodeContext) {}

// EnterSet_col_list is called when production set_col_list is entered.
func (s *BaseDmSqlParserListener) EnterSet_col_list(ctx *Set_col_listContext) {}

// ExitSet_col_list is called when production set_col_list is exited.
func (s *BaseDmSqlParserListener) ExitSet_col_list(ctx *Set_col_listContext) {}

// EnterUpdate_from_clause is called when production update_from_clause is entered.
func (s *BaseDmSqlParserListener) EnterUpdate_from_clause(ctx *Update_from_clauseContext) {}

// ExitUpdate_from_clause is called when production update_from_clause is exited.
func (s *BaseDmSqlParserListener) ExitUpdate_from_clause(ctx *Update_from_clauseContext) {}

// EnterMerge_into_stmt is called when production merge_into_stmt is entered.
func (s *BaseDmSqlParserListener) EnterMerge_into_stmt(ctx *Merge_into_stmtContext) {}

// ExitMerge_into_stmt is called when production merge_into_stmt is exited.
func (s *BaseDmSqlParserListener) ExitMerge_into_stmt(ctx *Merge_into_stmtContext) {}

// EnterMerge_into_stmt_body is called when production merge_into_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterMerge_into_stmt_body(ctx *Merge_into_stmt_bodyContext) {}

// ExitMerge_into_stmt_body is called when production merge_into_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitMerge_into_stmt_body(ctx *Merge_into_stmt_bodyContext) {}

// EnterMerge_into_sub_clause is called when production merge_into_sub_clause is entered.
func (s *BaseDmSqlParserListener) EnterMerge_into_sub_clause(ctx *Merge_into_sub_clauseContext) {}

// ExitMerge_into_sub_clause is called when production merge_into_sub_clause is exited.
func (s *BaseDmSqlParserListener) ExitMerge_into_sub_clause(ctx *Merge_into_sub_clauseContext) {}

// EnterMerge_update_clause is called when production merge_update_clause is entered.
func (s *BaseDmSqlParserListener) EnterMerge_update_clause(ctx *Merge_update_clauseContext) {}

// ExitMerge_update_clause is called when production merge_update_clause is exited.
func (s *BaseDmSqlParserListener) ExitMerge_update_clause(ctx *Merge_update_clauseContext) {}

// EnterMerge_insert_clause is called when production merge_insert_clause is entered.
func (s *BaseDmSqlParserListener) EnterMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// ExitMerge_insert_clause is called when production merge_insert_clause is exited.
func (s *BaseDmSqlParserListener) ExitMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// EnterCreate_profile_stmt is called when production create_profile_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_profile_stmt(ctx *Create_profile_stmtContext) {}

// ExitCreate_profile_stmt is called when production create_profile_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_profile_stmt(ctx *Create_profile_stmtContext) {}

// EnterAlter_profile_stmt is called when production alter_profile_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_profile_stmt(ctx *Alter_profile_stmtContext) {}

// ExitAlter_profile_stmt is called when production alter_profile_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_profile_stmt(ctx *Alter_profile_stmtContext) {}

// EnterCreate_user_stmt is called when production create_user_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_user_stmt(ctx *Create_user_stmtContext) {}

// ExitCreate_user_stmt is called when production create_user_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_user_stmt(ctx *Create_user_stmtContext) {}

// EnterDefault_ts_name is called when production default_ts_name is entered.
func (s *BaseDmSqlParserListener) EnterDefault_ts_name(ctx *Default_ts_nameContext) {}

// ExitDefault_ts_name is called when production default_ts_name is exited.
func (s *BaseDmSqlParserListener) ExitDefault_ts_name(ctx *Default_ts_nameContext) {}

// EnterDefault_ts_name_lst is called when production default_ts_name_lst is entered.
func (s *BaseDmSqlParserListener) EnterDefault_ts_name_lst(ctx *Default_ts_name_lstContext) {}

// ExitDefault_ts_name_lst is called when production default_ts_name_lst is exited.
func (s *BaseDmSqlParserListener) ExitDefault_ts_name_lst(ctx *Default_ts_name_lstContext) {}

// EnterDefault_ts_name_node is called when production default_ts_name_node is entered.
func (s *BaseDmSqlParserListener) EnterDefault_ts_name_node(ctx *Default_ts_name_nodeContext) {}

// ExitDefault_ts_name_node is called when production default_ts_name_node is exited.
func (s *BaseDmSqlParserListener) ExitDefault_ts_name_node(ctx *Default_ts_name_nodeContext) {}

// EnterDefault_idx_ts_name is called when production default_idx_ts_name is entered.
func (s *BaseDmSqlParserListener) EnterDefault_idx_ts_name(ctx *Default_idx_ts_nameContext) {}

// ExitDefault_idx_ts_name is called when production default_idx_ts_name is exited.
func (s *BaseDmSqlParserListener) ExitDefault_idx_ts_name(ctx *Default_idx_ts_nameContext) {}

// EnterDefault_ts_name_low is called when production default_ts_name_low is entered.
func (s *BaseDmSqlParserListener) EnterDefault_ts_name_low(ctx *Default_ts_name_lowContext) {}

// ExitDefault_ts_name_low is called when production default_ts_name_low is exited.
func (s *BaseDmSqlParserListener) ExitDefault_ts_name_low(ctx *Default_ts_name_lowContext) {}

// EnterTemp_ts_name is called when production temp_ts_name is entered.
func (s *BaseDmSqlParserListener) EnterTemp_ts_name(ctx *Temp_ts_nameContext) {}

// ExitTemp_ts_name is called when production temp_ts_name is exited.
func (s *BaseDmSqlParserListener) ExitTemp_ts_name(ctx *Temp_ts_nameContext) {}

// EnterDefault_ts_group_name_low is called when production default_ts_group_name_low is entered.
func (s *BaseDmSqlParserListener) EnterDefault_ts_group_name_low(ctx *Default_ts_group_name_lowContext) {
}

// ExitDefault_ts_group_name_low is called when production default_ts_group_name_low is exited.
func (s *BaseDmSqlParserListener) ExitDefault_ts_group_name_low(ctx *Default_ts_group_name_lowContext) {
}

// EnterOn_schema is called when production on_schema is entered.
func (s *BaseDmSqlParserListener) EnterOn_schema(ctx *On_schemaContext) {}

// ExitOn_schema is called when production on_schema is exited.
func (s *BaseDmSqlParserListener) ExitOn_schema(ctx *On_schemaContext) {}

// EnterReplace_old_pwd is called when production replace_old_pwd is entered.
func (s *BaseDmSqlParserListener) EnterReplace_old_pwd(ctx *Replace_old_pwdContext) {}

// ExitReplace_old_pwd is called when production replace_old_pwd is exited.
func (s *BaseDmSqlParserListener) ExitReplace_old_pwd(ctx *Replace_old_pwdContext) {}

// EnterAlter_user_stmt is called when production alter_user_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_user_stmt(ctx *Alter_user_stmtContext) {}

// ExitAlter_user_stmt is called when production alter_user_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_user_stmt(ctx *Alter_user_stmtContext) {}

// EnterUser_encrypt_options is called when production user_encrypt_options is entered.
func (s *BaseDmSqlParserListener) EnterUser_encrypt_options(ctx *User_encrypt_optionsContext) {}

// ExitUser_encrypt_options is called when production user_encrypt_options is exited.
func (s *BaseDmSqlParserListener) ExitUser_encrypt_options(ctx *User_encrypt_optionsContext) {}

// EnterUser_encrypt_option is called when production user_encrypt_option is entered.
func (s *BaseDmSqlParserListener) EnterUser_encrypt_option(ctx *User_encrypt_optionContext) {}

// ExitUser_encrypt_option is called when production user_encrypt_option is exited.
func (s *BaseDmSqlParserListener) ExitUser_encrypt_option(ctx *User_encrypt_optionContext) {}

// EnterAuthent_type_options is called when production authent_type_options is entered.
func (s *BaseDmSqlParserListener) EnterAuthent_type_options(ctx *Authent_type_optionsContext) {}

// ExitAuthent_type_options is called when production authent_type_options is exited.
func (s *BaseDmSqlParserListener) ExitAuthent_type_options(ctx *Authent_type_optionsContext) {}

// EnterHash_cipher_option is called when production hash_cipher_option is entered.
func (s *BaseDmSqlParserListener) EnterHash_cipher_option(ctx *Hash_cipher_optionContext) {}

// ExitHash_cipher_option is called when production hash_cipher_option is exited.
func (s *BaseDmSqlParserListener) ExitHash_cipher_option(ctx *Hash_cipher_optionContext) {}

// EnterAuthent_type is called when production authent_type is entered.
func (s *BaseDmSqlParserListener) EnterAuthent_type(ctx *Authent_typeContext) {}

// ExitAuthent_type is called when production authent_type is exited.
func (s *BaseDmSqlParserListener) ExitAuthent_type(ctx *Authent_typeContext) {}

// EnterForce_format is called when production force_format is entered.
func (s *BaseDmSqlParserListener) EnterForce_format(ctx *Force_formatContext) {}

// ExitForce_format is called when production force_format is exited.
func (s *BaseDmSqlParserListener) ExitForce_format(ctx *Force_formatContext) {}

// EnterAs is called when production as is entered.
func (s *BaseDmSqlParserListener) EnterAs(ctx *AsContext) {}

// ExitAs is called when production as is exited.
func (s *BaseDmSqlParserListener) ExitAs(ctx *AsContext) {}

// EnterPwd_policy is called when production pwd_policy is entered.
func (s *BaseDmSqlParserListener) EnterPwd_policy(ctx *Pwd_policyContext) {}

// ExitPwd_policy is called when production pwd_policy is exited.
func (s *BaseDmSqlParserListener) ExitPwd_policy(ctx *Pwd_policyContext) {}

// EnterAccount_lock is called when production account_lock is entered.
func (s *BaseDmSqlParserListener) EnterAccount_lock(ctx *Account_lockContext) {}

// ExitAccount_lock is called when production account_lock is exited.
func (s *BaseDmSqlParserListener) ExitAccount_lock(ctx *Account_lockContext) {}

// EnterRead_only_flag is called when production read_only_flag is entered.
func (s *BaseDmSqlParserListener) EnterRead_only_flag(ctx *Read_only_flagContext) {}

// ExitRead_only_flag is called when production read_only_flag is exited.
func (s *BaseDmSqlParserListener) ExitRead_only_flag(ctx *Read_only_flagContext) {}

// EnterRead_only_flag_not_null is called when production read_only_flag_not_null is entered.
func (s *BaseDmSqlParserListener) EnterRead_only_flag_not_null(ctx *Read_only_flag_not_nullContext) {}

// ExitRead_only_flag_not_null is called when production read_only_flag_not_null is exited.
func (s *BaseDmSqlParserListener) ExitRead_only_flag_not_null(ctx *Read_only_flag_not_nullContext) {}

// EnterResource is called when production resource is entered.
func (s *BaseDmSqlParserListener) EnterResource(ctx *ResourceContext) {}

// ExitResource is called when production resource is exited.
func (s *BaseDmSqlParserListener) ExitResource(ctx *ResourceContext) {}

// EnterExpire is called when production expire is entered.
func (s *BaseDmSqlParserListener) EnterExpire(ctx *ExpireContext) {}

// ExitExpire is called when production expire is exited.
func (s *BaseDmSqlParserListener) ExitExpire(ctx *ExpireContext) {}

// EnterResource_limit_options is called when production resource_limit_options is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit_options(ctx *Resource_limit_optionsContext) {}

// ExitResource_limit_options is called when production resource_limit_options is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit_options(ctx *Resource_limit_optionsContext) {}

// EnterResource_limit_list is called when production resource_limit_list is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit_list(ctx *Resource_limit_listContext) {}

// ExitResource_limit_list is called when production resource_limit_list is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit_list(ctx *Resource_limit_listContext) {}

// EnterResource_limit_list_with_comma is called when production resource_limit_list_with_comma is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit_list_with_comma(ctx *Resource_limit_list_with_commaContext) {
}

// ExitResource_limit_list_with_comma is called when production resource_limit_list_with_comma is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit_list_with_comma(ctx *Resource_limit_list_with_commaContext) {
}

// EnterResource_limit_list_with_empty is called when production resource_limit_list_with_empty is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit_list_with_empty(ctx *Resource_limit_list_with_emptyContext) {
}

// ExitResource_limit_list_with_empty is called when production resource_limit_list_with_empty is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit_list_with_empty(ctx *Resource_limit_list_with_emptyContext) {
}

// EnterResource_limit is called when production resource_limit is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit(ctx *Resource_limitContext) {}

// ExitResource_limit is called when production resource_limit is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit(ctx *Resource_limitContext) {}

// EnterResource_limit_value is called when production resource_limit_value is entered.
func (s *BaseDmSqlParserListener) EnterResource_limit_value(ctx *Resource_limit_valueContext) {}

// ExitResource_limit_value is called when production resource_limit_value is exited.
func (s *BaseDmSqlParserListener) ExitResource_limit_value(ctx *Resource_limit_valueContext) {}

// EnterCreate_audit_rule_stmt is called when production create_audit_rule_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_audit_rule_stmt(ctx *Create_audit_rule_stmtContext) {}

// ExitCreate_audit_rule_stmt is called when production create_audit_rule_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_audit_rule_stmt(ctx *Create_audit_rule_stmtContext) {}

// EnterRule_name is called when production rule_name is entered.
func (s *BaseDmSqlParserListener) EnterRule_name(ctx *Rule_nameContext) {}

// ExitRule_name is called when production rule_name is exited.
func (s *BaseDmSqlParserListener) ExitRule_name(ctx *Rule_nameContext) {}

// EnterAudit_rule_action is called when production audit_rule_action is entered.
func (s *BaseDmSqlParserListener) EnterAudit_rule_action(ctx *Audit_rule_actionContext) {}

// ExitAudit_rule_action is called when production audit_rule_action is exited.
func (s *BaseDmSqlParserListener) ExitAudit_rule_action(ctx *Audit_rule_actionContext) {}

// EnterBy_login_or_all is called when production by_login_or_all is entered.
func (s *BaseDmSqlParserListener) EnterBy_login_or_all(ctx *By_login_or_allContext) {}

// ExitBy_login_or_all is called when production by_login_or_all is exited.
func (s *BaseDmSqlParserListener) ExitBy_login_or_all(ctx *By_login_or_allContext) {}

// EnterAllow_ip_list is called when production allow_ip_list is entered.
func (s *BaseDmSqlParserListener) EnterAllow_ip_list(ctx *Allow_ip_listContext) {}

// ExitAllow_ip_list is called when production allow_ip_list is exited.
func (s *BaseDmSqlParserListener) ExitAllow_ip_list(ctx *Allow_ip_listContext) {}

// EnterNot_allow_ip_list is called when production not_allow_ip_list is entered.
func (s *BaseDmSqlParserListener) EnterNot_allow_ip_list(ctx *Not_allow_ip_listContext) {}

// ExitNot_allow_ip_list is called when production not_allow_ip_list is exited.
func (s *BaseDmSqlParserListener) ExitNot_allow_ip_list(ctx *Not_allow_ip_listContext) {}

// EnterIp_list is called when production ip_list is entered.
func (s *BaseDmSqlParserListener) EnterIp_list(ctx *Ip_listContext) {}

// ExitIp_list is called when production ip_list is exited.
func (s *BaseDmSqlParserListener) ExitIp_list(ctx *Ip_listContext) {}

// EnterAllow_dt_list is called when production allow_dt_list is entered.
func (s *BaseDmSqlParserListener) EnterAllow_dt_list(ctx *Allow_dt_listContext) {}

// ExitAllow_dt_list is called when production allow_dt_list is exited.
func (s *BaseDmSqlParserListener) ExitAllow_dt_list(ctx *Allow_dt_listContext) {}

// EnterNot_allow_dt_list is called when production not_allow_dt_list is entered.
func (s *BaseDmSqlParserListener) EnterNot_allow_dt_list(ctx *Not_allow_dt_listContext) {}

// ExitNot_allow_dt_list is called when production not_allow_dt_list is exited.
func (s *BaseDmSqlParserListener) ExitNot_allow_dt_list(ctx *Not_allow_dt_listContext) {}

// EnterDt_list is called when production dt_list is entered.
func (s *BaseDmSqlParserListener) EnterDt_list(ctx *Dt_listContext) {}

// ExitDt_list is called when production dt_list is exited.
func (s *BaseDmSqlParserListener) ExitDt_list(ctx *Dt_listContext) {}

// EnterDt is called when production dt is entered.
func (s *BaseDmSqlParserListener) EnterDt(ctx *DtContext) {}

// ExitDt is called when production dt is exited.
func (s *BaseDmSqlParserListener) ExitDt(ctx *DtContext) {}

// EnterOp_freq is called when production op_freq is entered.
func (s *BaseDmSqlParserListener) EnterOp_freq(ctx *Op_freqContext) {}

// ExitOp_freq is called when production op_freq is exited.
func (s *BaseDmSqlParserListener) ExitOp_freq(ctx *Op_freqContext) {}

// EnterQuota_val is called when production quota_val is entered.
func (s *BaseDmSqlParserListener) EnterQuota_val(ctx *Quota_valContext) {}

// ExitQuota_val is called when production quota_val is exited.
func (s *BaseDmSqlParserListener) ExitQuota_val(ctx *Quota_valContext) {}

// EnterQuota_ts_node is called when production quota_ts_node is entered.
func (s *BaseDmSqlParserListener) EnterQuota_ts_node(ctx *Quota_ts_nodeContext) {}

// ExitQuota_ts_node is called when production quota_ts_node is exited.
func (s *BaseDmSqlParserListener) ExitQuota_ts_node(ctx *Quota_ts_nodeContext) {}

// EnterQuota_ts_lst is called when production quota_ts_lst is entered.
func (s *BaseDmSqlParserListener) EnterQuota_ts_lst(ctx *Quota_ts_lstContext) {}

// ExitQuota_ts_lst is called when production quota_ts_lst is exited.
func (s *BaseDmSqlParserListener) ExitQuota_ts_lst(ctx *Quota_ts_lstContext) {}

// EnterQuota_ts is called when production quota_ts is entered.
func (s *BaseDmSqlParserListener) EnterQuota_ts(ctx *Quota_tsContext) {}

// ExitQuota_ts is called when production quota_ts is exited.
func (s *BaseDmSqlParserListener) ExitQuota_ts(ctx *Quota_tsContext) {}

// EnterCreate_role_stmt is called when production create_role_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_role_stmt(ctx *Create_role_stmtContext) {}

// ExitCreate_role_stmt is called when production create_role_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_role_stmt(ctx *Create_role_stmtContext) {}

// EnterCreate_dblink_stmt is called when production create_dblink_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_dblink_stmt(ctx *Create_dblink_stmtContext) {}

// ExitCreate_dblink_stmt is called when production create_dblink_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_dblink_stmt(ctx *Create_dblink_stmtContext) {}

// EnterDb_type_str is called when production db_type_str is entered.
func (s *BaseDmSqlParserListener) EnterDb_type_str(ctx *Db_type_strContext) {}

// ExitDb_type_str is called when production db_type_str is exited.
func (s *BaseDmSqlParserListener) ExitDb_type_str(ctx *Db_type_strContext) {}

// EnterDblink_option_lst_options is called when production dblink_option_lst_options is entered.
func (s *BaseDmSqlParserListener) EnterDblink_option_lst_options(ctx *Dblink_option_lst_optionsContext) {
}

// ExitDblink_option_lst_options is called when production dblink_option_lst_options is exited.
func (s *BaseDmSqlParserListener) ExitDblink_option_lst_options(ctx *Dblink_option_lst_optionsContext) {
}

// EnterDblink_option_lst is called when production dblink_option_lst is entered.
func (s *BaseDmSqlParserListener) EnterDblink_option_lst(ctx *Dblink_option_lstContext) {}

// ExitDblink_option_lst is called when production dblink_option_lst is exited.
func (s *BaseDmSqlParserListener) ExitDblink_option_lst(ctx *Dblink_option_lstContext) {}

// EnterDblink_option is called when production dblink_option is entered.
func (s *BaseDmSqlParserListener) EnterDblink_option(ctx *Dblink_optionContext) {}

// ExitDblink_option is called when production dblink_option is exited.
func (s *BaseDmSqlParserListener) ExitDblink_option(ctx *Dblink_optionContext) {}

// EnterCreate_synonym_stmt is called when production create_synonym_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_synonym_stmt(ctx *Create_synonym_stmtContext) {}

// ExitCreate_synonym_stmt is called when production create_synonym_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_synonym_stmt(ctx *Create_synonym_stmtContext) {}

// EnterFull_synonym_name is called when production full_synonym_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_synonym_name(ctx *Full_synonym_nameContext) {}

// ExitFull_synonym_name is called when production full_synonym_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_synonym_name(ctx *Full_synonym_nameContext) {}

// EnterFull_obj_name is called when production full_obj_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_obj_name(ctx *Full_obj_nameContext) {}

// ExitFull_obj_name is called when production full_obj_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_obj_name(ctx *Full_obj_nameContext) {}

// EnterCreate_domain_stmt is called when production create_domain_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_domain_stmt(ctx *Create_domain_stmtContext) {}

// ExitCreate_domain_stmt is called when production create_domain_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_domain_stmt(ctx *Create_domain_stmtContext) {}

// EnterDomain_default_option is called when production domain_default_option is entered.
func (s *BaseDmSqlParserListener) EnterDomain_default_option(ctx *Domain_default_optionContext) {}

// ExitDomain_default_option is called when production domain_default_option is exited.
func (s *BaseDmSqlParserListener) ExitDomain_default_option(ctx *Domain_default_optionContext) {}

// EnterDomain_constraints_option is called when production domain_constraints_option is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraints_option(ctx *Domain_constraints_optionContext) {
}

// ExitDomain_constraints_option is called when production domain_constraints_option is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraints_option(ctx *Domain_constraints_optionContext) {
}

// EnterDomain_constraints_def is called when production domain_constraints_def is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraints_def(ctx *Domain_constraints_defContext) {}

// ExitDomain_constraints_def is called when production domain_constraints_def is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraints_def(ctx *Domain_constraints_defContext) {}

// EnterDomain_constraints is called when production domain_constraints is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraints(ctx *Domain_constraintsContext) {}

// ExitDomain_constraints is called when production domain_constraints is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraints(ctx *Domain_constraintsContext) {}

// EnterDomain_constraint is called when production domain_constraint is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraint(ctx *Domain_constraintContext) {}

// ExitDomain_constraint is called when production domain_constraint is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraint(ctx *Domain_constraintContext) {}

// EnterDomain_constraint_name_def_options is called when production domain_constraint_name_def_options is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraint_name_def_options(ctx *Domain_constraint_name_def_optionsContext) {
}

// ExitDomain_constraint_name_def_options is called when production domain_constraint_name_def_options is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraint_name_def_options(ctx *Domain_constraint_name_def_optionsContext) {
}

// EnterDomain_constraint_name_def is called when production domain_constraint_name_def is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraint_name_def(ctx *Domain_constraint_name_defContext) {
}

// ExitDomain_constraint_name_def is called when production domain_constraint_name_def is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraint_name_def(ctx *Domain_constraint_name_defContext) {
}

// EnterDomain_constraint_name is called when production domain_constraint_name is entered.
func (s *BaseDmSqlParserListener) EnterDomain_constraint_name(ctx *Domain_constraint_nameContext) {}

// ExitDomain_constraint_name is called when production domain_constraint_name is exited.
func (s *BaseDmSqlParserListener) ExitDomain_constraint_name(ctx *Domain_constraint_nameContext) {}

// EnterCreate_character_set_stmt is called when production create_character_set_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_character_set_stmt(ctx *Create_character_set_stmtContext) {
}

// ExitCreate_character_set_stmt is called when production create_character_set_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_character_set_stmt(ctx *Create_character_set_stmtContext) {
}

// EnterCharacter_set_source is called when production character_set_source is entered.
func (s *BaseDmSqlParserListener) EnterCharacter_set_source(ctx *Character_set_sourceContext) {}

// ExitCharacter_set_source is called when production character_set_source is exited.
func (s *BaseDmSqlParserListener) ExitCharacter_set_source(ctx *Character_set_sourceContext) {}

// EnterExisting_character_set_name is called when production existing_character_set_name is entered.
func (s *BaseDmSqlParserListener) EnterExisting_character_set_name(ctx *Existing_character_set_nameContext) {
}

// ExitExisting_character_set_name is called when production existing_character_set_name is exited.
func (s *BaseDmSqlParserListener) ExitExisting_character_set_name(ctx *Existing_character_set_nameContext) {
}

// EnterCharacter_set_name is called when production character_set_name is entered.
func (s *BaseDmSqlParserListener) EnterCharacter_set_name(ctx *Character_set_nameContext) {}

// ExitCharacter_set_name is called when production character_set_name is exited.
func (s *BaseDmSqlParserListener) ExitCharacter_set_name(ctx *Character_set_nameContext) {}

// EnterCollate_clause_option is called when production collate_clause_option is entered.
func (s *BaseDmSqlParserListener) EnterCollate_clause_option(ctx *Collate_clause_optionContext) {}

// ExitCollate_clause_option is called when production collate_clause_option is exited.
func (s *BaseDmSqlParserListener) ExitCollate_clause_option(ctx *Collate_clause_optionContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseDmSqlParserListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseDmSqlParserListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterCreate_collation_stmt is called when production create_collation_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_collation_stmt(ctx *Create_collation_stmtContext) {}

// ExitCreate_collation_stmt is called when production create_collation_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_collation_stmt(ctx *Create_collation_stmtContext) {}

// EnterExisting_collation_name is called when production existing_collation_name is entered.
func (s *BaseDmSqlParserListener) EnterExisting_collation_name(ctx *Existing_collation_nameContext) {}

// ExitExisting_collation_name is called when production existing_collation_name is exited.
func (s *BaseDmSqlParserListener) ExitExisting_collation_name(ctx *Existing_collation_nameContext) {}

// EnterPad_option is called when production pad_option is entered.
func (s *BaseDmSqlParserListener) EnterPad_option(ctx *Pad_optionContext) {}

// ExitPad_option is called when production pad_option is exited.
func (s *BaseDmSqlParserListener) ExitPad_option(ctx *Pad_optionContext) {}

// EnterCreate_sequence_stmt is called when production create_sequence_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_sequence_stmt(ctx *Create_sequence_stmtContext) {}

// ExitCreate_sequence_stmt is called when production create_sequence_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_sequence_stmt(ctx *Create_sequence_stmtContext) {}

// EnterSequence_option_list_options is called when production sequence_option_list_options is entered.
func (s *BaseDmSqlParserListener) EnterSequence_option_list_options(ctx *Sequence_option_list_optionsContext) {
}

// ExitSequence_option_list_options is called when production sequence_option_list_options is exited.
func (s *BaseDmSqlParserListener) ExitSequence_option_list_options(ctx *Sequence_option_list_optionsContext) {
}

// EnterSequence_option_list is called when production sequence_option_list is entered.
func (s *BaseDmSqlParserListener) EnterSequence_option_list(ctx *Sequence_option_listContext) {}

// ExitSequence_option_list is called when production sequence_option_list is exited.
func (s *BaseDmSqlParserListener) ExitSequence_option_list(ctx *Sequence_option_listContext) {}

// EnterSequence_option is called when production sequence_option is entered.
func (s *BaseDmSqlParserListener) EnterSequence_option(ctx *Sequence_optionContext) {}

// ExitSequence_option is called when production sequence_option is exited.
func (s *BaseDmSqlParserListener) ExitSequence_option(ctx *Sequence_optionContext) {}

// EnterSequence_name is called when production sequence_name is entered.
func (s *BaseDmSqlParserListener) EnterSequence_name(ctx *Sequence_nameContext) {}

// ExitSequence_name is called when production sequence_name is exited.
func (s *BaseDmSqlParserListener) ExitSequence_name(ctx *Sequence_nameContext) {}

// EnterIncrement_option is called when production increment_option is entered.
func (s *BaseDmSqlParserListener) EnterIncrement_option(ctx *Increment_optionContext) {}

// ExitIncrement_option is called when production increment_option is exited.
func (s *BaseDmSqlParserListener) ExitIncrement_option(ctx *Increment_optionContext) {}

// EnterStart_option is called when production start_option is entered.
func (s *BaseDmSqlParserListener) EnterStart_option(ctx *Start_optionContext) {}

// ExitStart_option is called when production start_option is exited.
func (s *BaseDmSqlParserListener) ExitStart_option(ctx *Start_optionContext) {}

// EnterCurrent_option is called when production current_option is entered.
func (s *BaseDmSqlParserListener) EnterCurrent_option(ctx *Current_optionContext) {}

// ExitCurrent_option is called when production current_option is exited.
func (s *BaseDmSqlParserListener) ExitCurrent_option(ctx *Current_optionContext) {}

// EnterMaxvalue_option is called when production maxvalue_option is entered.
func (s *BaseDmSqlParserListener) EnterMaxvalue_option(ctx *Maxvalue_optionContext) {}

// ExitMaxvalue_option is called when production maxvalue_option is exited.
func (s *BaseDmSqlParserListener) ExitMaxvalue_option(ctx *Maxvalue_optionContext) {}

// EnterMinvalue_option is called when production minvalue_option is entered.
func (s *BaseDmSqlParserListener) EnterMinvalue_option(ctx *Minvalue_optionContext) {}

// ExitMinvalue_option is called when production minvalue_option is exited.
func (s *BaseDmSqlParserListener) ExitMinvalue_option(ctx *Minvalue_optionContext) {}

// EnterCycle_option is called when production cycle_option is entered.
func (s *BaseDmSqlParserListener) EnterCycle_option(ctx *Cycle_optionContext) {}

// ExitCycle_option is called when production cycle_option is exited.
func (s *BaseDmSqlParserListener) ExitCycle_option(ctx *Cycle_optionContext) {}

// EnterCache_option is called when production cache_option is entered.
func (s *BaseDmSqlParserListener) EnterCache_option(ctx *Cache_optionContext) {}

// ExitCache_option is called when production cache_option is exited.
func (s *BaseDmSqlParserListener) ExitCache_option(ctx *Cache_optionContext) {}

// EnterOrder_option is called when production order_option is entered.
func (s *BaseDmSqlParserListener) EnterOrder_option(ctx *Order_optionContext) {}

// ExitOrder_option is called when production order_option is exited.
func (s *BaseDmSqlParserListener) ExitOrder_option(ctx *Order_optionContext) {}

// EnterSeq_local_option is called when production seq_local_option is entered.
func (s *BaseDmSqlParserListener) EnterSeq_local_option(ctx *Seq_local_optionContext) {}

// ExitSeq_local_option is called when production seq_local_option is exited.
func (s *BaseDmSqlParserListener) ExitSeq_local_option(ctx *Seq_local_optionContext) {}

// EnterWhenever_stmt_options is called when production whenever_stmt_options is entered.
func (s *BaseDmSqlParserListener) EnterWhenever_stmt_options(ctx *Whenever_stmt_optionsContext) {}

// ExitWhenever_stmt_options is called when production whenever_stmt_options is exited.
func (s *BaseDmSqlParserListener) ExitWhenever_stmt_options(ctx *Whenever_stmt_optionsContext) {}

// EnterWhenever_stmt is called when production whenever_stmt is entered.
func (s *BaseDmSqlParserListener) EnterWhenever_stmt(ctx *Whenever_stmtContext) {}

// ExitWhenever_stmt is called when production whenever_stmt is exited.
func (s *BaseDmSqlParserListener) ExitWhenever_stmt(ctx *Whenever_stmtContext) {}

// EnterGrant_stmt is called when production grant_stmt is entered.
func (s *BaseDmSqlParserListener) EnterGrant_stmt(ctx *Grant_stmtContext) {}

// ExitGrant_stmt is called when production grant_stmt is exited.
func (s *BaseDmSqlParserListener) ExitGrant_stmt(ctx *Grant_stmtContext) {}

// EnterGrant_tag is called when production grant_tag is entered.
func (s *BaseDmSqlParserListener) EnterGrant_tag(ctx *Grant_tagContext) {}

// ExitGrant_tag is called when production grant_tag is exited.
func (s *BaseDmSqlParserListener) ExitGrant_tag(ctx *Grant_tagContext) {}

// EnterGrant_stmt_body is called when production grant_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterGrant_stmt_body(ctx *Grant_stmt_bodyContext) {}

// ExitGrant_stmt_body is called when production grant_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitGrant_stmt_body(ctx *Grant_stmt_bodyContext) {}

// EnterRevoke_stmt is called when production revoke_stmt is entered.
func (s *BaseDmSqlParserListener) EnterRevoke_stmt(ctx *Revoke_stmtContext) {}

// ExitRevoke_stmt is called when production revoke_stmt is exited.
func (s *BaseDmSqlParserListener) ExitRevoke_stmt(ctx *Revoke_stmtContext) {}

// EnterRevoke_stmt_body is called when production revoke_stmt_body is entered.
func (s *BaseDmSqlParserListener) EnterRevoke_stmt_body(ctx *Revoke_stmt_bodyContext) {}

// ExitRevoke_stmt_body is called when production revoke_stmt_body is exited.
func (s *BaseDmSqlParserListener) ExitRevoke_stmt_body(ctx *Revoke_stmt_bodyContext) {}

// EnterGrant_privs is called when production grant_privs is entered.
func (s *BaseDmSqlParserListener) EnterGrant_privs(ctx *Grant_privsContext) {}

// ExitGrant_privs is called when production grant_privs is exited.
func (s *BaseDmSqlParserListener) ExitGrant_privs(ctx *Grant_privsContext) {}

// EnterGrant_priv_list is called when production grant_priv_list is entered.
func (s *BaseDmSqlParserListener) EnterGrant_priv_list(ctx *Grant_priv_listContext) {}

// ExitGrant_priv_list is called when production grant_priv_list is exited.
func (s *BaseDmSqlParserListener) ExitGrant_priv_list(ctx *Grant_priv_listContext) {}

// EnterGrant_priv_off is called when production grant_priv_off is entered.
func (s *BaseDmSqlParserListener) EnterGrant_priv_off(ctx *Grant_priv_offContext) {}

// ExitGrant_priv_off is called when production grant_priv_off is exited.
func (s *BaseDmSqlParserListener) ExitGrant_priv_off(ctx *Grant_priv_offContext) {}

// EnterGrant_priv is called when production grant_priv is entered.
func (s *BaseDmSqlParserListener) EnterGrant_priv(ctx *Grant_privContext) {}

// ExitGrant_priv is called when production grant_priv is exited.
func (s *BaseDmSqlParserListener) ExitGrant_priv(ctx *Grant_privContext) {}

// EnterRevoke_action is called when production revoke_action is entered.
func (s *BaseDmSqlParserListener) EnterRevoke_action(ctx *Revoke_actionContext) {}

// ExitRevoke_action is called when production revoke_action is exited.
func (s *BaseDmSqlParserListener) ExitRevoke_action(ctx *Revoke_actionContext) {}

// EnterDb_priv_list is called when production db_priv_list is entered.
func (s *BaseDmSqlParserListener) EnterDb_priv_list(ctx *Db_priv_listContext) {}

// ExitDb_priv_list is called when production db_priv_list is exited.
func (s *BaseDmSqlParserListener) ExitDb_priv_list(ctx *Db_priv_listContext) {}

// EnterDb_priv is called when production db_priv is entered.
func (s *BaseDmSqlParserListener) EnterDb_priv(ctx *Db_privContext) {}

// ExitDb_priv is called when production db_priv is exited.
func (s *BaseDmSqlParserListener) ExitDb_priv(ctx *Db_privContext) {}

// EnterBy_grantor is called when production by_grantor is entered.
func (s *BaseDmSqlParserListener) EnterBy_grantor(ctx *By_grantorContext) {}

// ExitBy_grantor is called when production by_grantor is exited.
func (s *BaseDmSqlParserListener) ExitBy_grantor(ctx *By_grantorContext) {}

// EnterGrantees is called when production grantees is entered.
func (s *BaseDmSqlParserListener) EnterGrantees(ctx *GranteesContext) {}

// ExitGrantees is called when production grantees is exited.
func (s *BaseDmSqlParserListener) ExitGrantees(ctx *GranteesContext) {}

// EnterCreate_schema_stmt is called when production create_schema_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_schema_stmt(ctx *Create_schema_stmtContext) {}

// ExitCreate_schema_stmt is called when production create_schema_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_schema_stmt(ctx *Create_schema_stmtContext) {}

// EnterOprt_arg is called when production oprt_arg is entered.
func (s *BaseDmSqlParserListener) EnterOprt_arg(ctx *Oprt_argContext) {}

// ExitOprt_arg is called when production oprt_arg is exited.
func (s *BaseDmSqlParserListener) ExitOprt_arg(ctx *Oprt_argContext) {}

// EnterOprt_arg_lst is called when production oprt_arg_lst is entered.
func (s *BaseDmSqlParserListener) EnterOprt_arg_lst(ctx *Oprt_arg_lstContext) {}

// ExitOprt_arg_lst is called when production oprt_arg_lst is exited.
func (s *BaseDmSqlParserListener) ExitOprt_arg_lst(ctx *Oprt_arg_lstContext) {}

// EnterCreate_operator_stmt is called when production create_operator_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_operator_stmt(ctx *Create_operator_stmtContext) {}

// ExitCreate_operator_stmt is called when production create_operator_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_operator_stmt(ctx *Create_operator_stmtContext) {}

// EnterDrop_operator_stmt is called when production drop_operator_stmt is entered.
func (s *BaseDmSqlParserListener) EnterDrop_operator_stmt(ctx *Drop_operator_stmtContext) {}

// ExitDrop_operator_stmt is called when production drop_operator_stmt is exited.
func (s *BaseDmSqlParserListener) ExitDrop_operator_stmt(ctx *Drop_operator_stmtContext) {}

// EnterGrant_and_ddl is called when production grant_and_ddl is entered.
func (s *BaseDmSqlParserListener) EnterGrant_and_ddl(ctx *Grant_and_ddlContext) {}

// ExitGrant_and_ddl is called when production grant_and_ddl is exited.
func (s *BaseDmSqlParserListener) ExitGrant_and_ddl(ctx *Grant_and_ddlContext) {}

// EnterTop_exp is called when production top_exp is entered.
func (s *BaseDmSqlParserListener) EnterTop_exp(ctx *Top_expContext) {}

// ExitTop_exp is called when production top_exp is exited.
func (s *BaseDmSqlParserListener) ExitTop_exp(ctx *Top_expContext) {}

// EnterU_oprt is called when production u_oprt is entered.
func (s *BaseDmSqlParserListener) EnterU_oprt(ctx *U_oprtContext) {}

// ExitU_oprt is called when production u_oprt is exited.
func (s *BaseDmSqlParserListener) ExitU_oprt(ctx *U_oprtContext) {}

// EnterQualified_u_oprt is called when production qualified_u_oprt is entered.
func (s *BaseDmSqlParserListener) EnterQualified_u_oprt(ctx *Qualified_u_oprtContext) {}

// ExitQualified_u_oprt is called when production qualified_u_oprt is exited.
func (s *BaseDmSqlParserListener) ExitQualified_u_oprt(ctx *Qualified_u_oprtContext) {}

// EnterExp_u_oprt is called when production exp_u_oprt is entered.
func (s *BaseDmSqlParserListener) EnterExp_u_oprt(ctx *Exp_u_oprtContext) {}

// ExitExp_u_oprt is called when production exp_u_oprt is exited.
func (s *BaseDmSqlParserListener) ExitExp_u_oprt(ctx *Exp_u_oprtContext) {}

// EnterRaw_exp is called when production raw_exp is entered.
func (s *BaseDmSqlParserListener) EnterRaw_exp(ctx *Raw_expContext) {}

// ExitRaw_exp is called when production raw_exp is exited.
func (s *BaseDmSqlParserListener) ExitRaw_exp(ctx *Raw_expContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseDmSqlParserListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseDmSqlParserListener) ExitExp(ctx *ExpContext) {}

// EnterFrom_first_last_option is called when production from_first_last_option is entered.
func (s *BaseDmSqlParserListener) EnterFrom_first_last_option(ctx *From_first_last_optionContext) {}

// ExitFrom_first_last_option is called when production from_first_last_option is exited.
func (s *BaseDmSqlParserListener) ExitFrom_first_last_option(ctx *From_first_last_optionContext) {}

// EnterAfun_arg_lst is called when production afun_arg_lst is entered.
func (s *BaseDmSqlParserListener) EnterAfun_arg_lst(ctx *Afun_arg_lstContext) {}

// ExitAfun_arg_lst is called when production afun_arg_lst is exited.
func (s *BaseDmSqlParserListener) ExitAfun_arg_lst(ctx *Afun_arg_lstContext) {}

// EnterAfun_arg_lst_low is called when production afun_arg_lst_low is entered.
func (s *BaseDmSqlParserListener) EnterAfun_arg_lst_low(ctx *Afun_arg_lst_lowContext) {}

// ExitAfun_arg_lst_low is called when production afun_arg_lst_low is exited.
func (s *BaseDmSqlParserListener) ExitAfun_arg_lst_low(ctx *Afun_arg_lst_lowContext) {}

// EnterIn_value_exp is called when production in_value_exp is entered.
func (s *BaseDmSqlParserListener) EnterIn_value_exp(ctx *In_value_expContext) {}

// ExitIn_value_exp is called when production in_value_exp is exited.
func (s *BaseDmSqlParserListener) ExitIn_value_exp(ctx *In_value_expContext) {}

// EnterAfun_partition_by is called when production afun_partition_by is entered.
func (s *BaseDmSqlParserListener) EnterAfun_partition_by(ctx *Afun_partition_byContext) {}

// ExitAfun_partition_by is called when production afun_partition_by is exited.
func (s *BaseDmSqlParserListener) ExitAfun_partition_by(ctx *Afun_partition_byContext) {}

// EnterAfun_windowing is called when production afun_windowing is entered.
func (s *BaseDmSqlParserListener) EnterAfun_windowing(ctx *Afun_windowingContext) {}

// ExitAfun_windowing is called when production afun_windowing is exited.
func (s *BaseDmSqlParserListener) ExitAfun_windowing(ctx *Afun_windowingContext) {}

// EnterAfun_windowing_type is called when production afun_windowing_type is entered.
func (s *BaseDmSqlParserListener) EnterAfun_windowing_type(ctx *Afun_windowing_typeContext) {}

// ExitAfun_windowing_type is called when production afun_windowing_type is exited.
func (s *BaseDmSqlParserListener) ExitAfun_windowing_type(ctx *Afun_windowing_typeContext) {}

// EnterAfun_range_clause is called when production afun_range_clause is entered.
func (s *BaseDmSqlParserListener) EnterAfun_range_clause(ctx *Afun_range_clauseContext) {}

// ExitAfun_range_clause is called when production afun_range_clause is exited.
func (s *BaseDmSqlParserListener) ExitAfun_range_clause(ctx *Afun_range_clauseContext) {}

// EnterPexp is called when production pexp is entered.
func (s *BaseDmSqlParserListener) EnterPexp(ctx *PexpContext) {}

// ExitPexp is called when production pexp is exited.
func (s *BaseDmSqlParserListener) ExitPexp(ctx *PexpContext) {}

// EnterPexp_pfx is called when production pexp_pfx is entered.
func (s *BaseDmSqlParserListener) EnterPexp_pfx(ctx *Pexp_pfxContext) {}

// ExitPexp_pfx is called when production pexp_pfx is exited.
func (s *BaseDmSqlParserListener) ExitPexp_pfx(ctx *Pexp_pfxContext) {}

// EnterPexp_cast is called when production pexp_cast is entered.
func (s *BaseDmSqlParserListener) EnterPexp_cast(ctx *Pexp_castContext) {}

// ExitPexp_cast is called when production pexp_cast is exited.
func (s *BaseDmSqlParserListener) ExitPexp_cast(ctx *Pexp_castContext) {}

// EnterPexp_b is called when production pexp_b is entered.
func (s *BaseDmSqlParserListener) EnterPexp_b(ctx *Pexp_bContext) {}

// ExitPexp_b is called when production pexp_b is exited.
func (s *BaseDmSqlParserListener) ExitPexp_b(ctx *Pexp_bContext) {}

// EnterPexp_a is called when production pexp_a is entered.
func (s *BaseDmSqlParserListener) EnterPexp_a(ctx *Pexp_aContext) {}

// ExitPexp_a is called when production pexp_a is exited.
func (s *BaseDmSqlParserListener) ExitPexp_a(ctx *Pexp_aContext) {}

// EnterPexp_c is called when production pexp_c is entered.
func (s *BaseDmSqlParserListener) EnterPexp_c(ctx *Pexp_cContext) {}

// ExitPexp_c is called when production pexp_c is exited.
func (s *BaseDmSqlParserListener) ExitPexp_c(ctx *Pexp_cContext) {}

// EnterPexp_c_insert is called when production pexp_c_insert is entered.
func (s *BaseDmSqlParserListener) EnterPexp_c_insert(ctx *Pexp_c_insertContext) {}

// ExitPexp_c_insert is called when production pexp_c_insert is exited.
func (s *BaseDmSqlParserListener) ExitPexp_c_insert(ctx *Pexp_c_insertContext) {}

// EnterPexp_d is called when production pexp_d is entered.
func (s *BaseDmSqlParserListener) EnterPexp_d(ctx *Pexp_dContext) {}

// ExitPexp_d is called when production pexp_d is exited.
func (s *BaseDmSqlParserListener) ExitPexp_d(ctx *Pexp_dContext) {}

// EnterPexp_e is called when production pexp_e is entered.
func (s *BaseDmSqlParserListener) EnterPexp_e(ctx *Pexp_eContext) {}

// ExitPexp_e is called when production pexp_e is exited.
func (s *BaseDmSqlParserListener) ExitPexp_e(ctx *Pexp_eContext) {}

// EnterPexp_pfx2 is called when production pexp_pfx2 is entered.
func (s *BaseDmSqlParserListener) EnterPexp_pfx2(ctx *Pexp_pfx2Context) {}

// ExitPexp_pfx2 is called when production pexp_pfx2 is exited.
func (s *BaseDmSqlParserListener) ExitPexp_pfx2(ctx *Pexp_pfx2Context) {}

// EnterMember2 is called when production member2 is entered.
func (s *BaseDmSqlParserListener) EnterMember2(ctx *Member2Context) {}

// ExitMember2 is called when production member2 is exited.
func (s *BaseDmSqlParserListener) ExitMember2(ctx *Member2Context) {}

// EnterPexp_c2_insert is called when production pexp_c2_insert is entered.
func (s *BaseDmSqlParserListener) EnterPexp_c2_insert(ctx *Pexp_c2_insertContext) {}

// ExitPexp_c2_insert is called when production pexp_c2_insert is exited.
func (s *BaseDmSqlParserListener) ExitPexp_c2_insert(ctx *Pexp_c2_insertContext) {}

// EnterMember_access2 is called when production member_access2 is entered.
func (s *BaseDmSqlParserListener) EnterMember_access2(ctx *Member_access2Context) {}

// ExitMember_access2 is called when production member_access2 is exited.
func (s *BaseDmSqlParserListener) ExitMember_access2(ctx *Member_access2Context) {}

// EnterInvocation_expression2 is called when production invocation_expression2 is entered.
func (s *BaseDmSqlParserListener) EnterInvocation_expression2(ctx *Invocation_expression2Context) {}

// ExitInvocation_expression2 is called when production invocation_expression2 is exited.
func (s *BaseDmSqlParserListener) ExitInvocation_expression2(ctx *Invocation_expression2Context) {}

// EnterMember is called when production member is entered.
func (s *BaseDmSqlParserListener) EnterMember(ctx *MemberContext) {}

// ExitMember is called when production member is exited.
func (s *BaseDmSqlParserListener) ExitMember(ctx *MemberContext) {}

// EnterKey is called when production key is entered.
func (s *BaseDmSqlParserListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseDmSqlParserListener) ExitKey(ctx *KeyContext) {}

// EnterMember_access is called when production member_access is entered.
func (s *BaseDmSqlParserListener) EnterMember_access(ctx *Member_accessContext) {}

// ExitMember_access is called when production member_access is exited.
func (s *BaseDmSqlParserListener) ExitMember_access(ctx *Member_accessContext) {}

// EnterInvocation_expression is called when production invocation_expression is entered.
func (s *BaseDmSqlParserListener) EnterInvocation_expression(ctx *Invocation_expressionContext) {}

// ExitInvocation_expression is called when production invocation_expression is exited.
func (s *BaseDmSqlParserListener) ExitInvocation_expression(ctx *Invocation_expressionContext) {}

// EnterInvocation_expression_low is called when production invocation_expression_low is entered.
func (s *BaseDmSqlParserListener) EnterInvocation_expression_low(ctx *Invocation_expression_lowContext) {
}

// ExitInvocation_expression_low is called when production invocation_expression_low is exited.
func (s *BaseDmSqlParserListener) ExitInvocation_expression_low(ctx *Invocation_expression_lowContext) {
}

// EnterXmlagg_inv_expression is called when production xmlagg_inv_expression is entered.
func (s *BaseDmSqlParserListener) EnterXmlagg_inv_expression(ctx *Xmlagg_inv_expressionContext) {}

// ExitXmlagg_inv_expression is called when production xmlagg_inv_expression is exited.
func (s *BaseDmSqlParserListener) ExitXmlagg_inv_expression(ctx *Xmlagg_inv_expressionContext) {}

// EnterXmlfun_inv_expression is called when production xmlfun_inv_expression is entered.
func (s *BaseDmSqlParserListener) EnterXmlfun_inv_expression(ctx *Xmlfun_inv_expressionContext) {}

// ExitXmlfun_inv_expression is called when production xmlfun_inv_expression is exited.
func (s *BaseDmSqlParserListener) ExitXmlfun_inv_expression(ctx *Xmlfun_inv_expressionContext) {}

// EnterXmlele_name is called when production xmlele_name is entered.
func (s *BaseDmSqlParserListener) EnterXmlele_name(ctx *Xmlele_nameContext) {}

// ExitXmlele_name is called when production xmlele_name is exited.
func (s *BaseDmSqlParserListener) ExitXmlele_name(ctx *Xmlele_nameContext) {}

// EnterXmlele_sub_lst is called when production xmlele_sub_lst is entered.
func (s *BaseDmSqlParserListener) EnterXmlele_sub_lst(ctx *Xmlele_sub_lstContext) {}

// ExitXmlele_sub_lst is called when production xmlele_sub_lst is exited.
func (s *BaseDmSqlParserListener) ExitXmlele_sub_lst(ctx *Xmlele_sub_lstContext) {}

// EnterXmlattr_lst is called when production xmlattr_lst is entered.
func (s *BaseDmSqlParserListener) EnterXmlattr_lst(ctx *Xmlattr_lstContext) {}

// ExitXmlattr_lst is called when production xmlattr_lst is exited.
func (s *BaseDmSqlParserListener) ExitXmlattr_lst(ctx *Xmlattr_lstContext) {}

// EnterXmlattr is called when production xmlattr is entered.
func (s *BaseDmSqlParserListener) EnterXmlattr(ctx *XmlattrContext) {}

// ExitXmlattr is called when production xmlattr is exited.
func (s *BaseDmSqlParserListener) ExitXmlattr(ctx *XmlattrContext) {}

// EnterXmlval_lst is called when production xmlval_lst is entered.
func (s *BaseDmSqlParserListener) EnterXmlval_lst(ctx *Xmlval_lstContext) {}

// ExitXmlval_lst is called when production xmlval_lst is exited.
func (s *BaseDmSqlParserListener) ExitXmlval_lst(ctx *Xmlval_lstContext) {}

// EnterKeep_clause is called when production keep_clause is entered.
func (s *BaseDmSqlParserListener) EnterKeep_clause(ctx *Keep_clauseContext) {}

// ExitKeep_clause is called when production keep_clause is exited.
func (s *BaseDmSqlParserListener) ExitKeep_clause(ctx *Keep_clauseContext) {}

// EnterWithin_clause is called when production within_clause is entered.
func (s *BaseDmSqlParserListener) EnterWithin_clause(ctx *Within_clauseContext) {}

// ExitWithin_clause is called when production within_clause is exited.
func (s *BaseDmSqlParserListener) ExitWithin_clause(ctx *Within_clauseContext) {}

// EnterTypeof_expression is called when production typeof_expression is entered.
func (s *BaseDmSqlParserListener) EnterTypeof_expression(ctx *Typeof_expressionContext) {}

// ExitTypeof_expression is called when production typeof_expression is exited.
func (s *BaseDmSqlParserListener) ExitTypeof_expression(ctx *Typeof_expressionContext) {}

// EnterNew_obj_expression is called when production new_obj_expression is entered.
func (s *BaseDmSqlParserListener) EnterNew_obj_expression(ctx *New_obj_expressionContext) {}

// ExitNew_obj_expression is called when production new_obj_expression is exited.
func (s *BaseDmSqlParserListener) ExitNew_obj_expression(ctx *New_obj_expressionContext) {}

// EnterNew_arr_expression is called when production new_arr_expression is entered.
func (s *BaseDmSqlParserListener) EnterNew_arr_expression(ctx *New_arr_expressionContext) {}

// ExitNew_arr_expression is called when production new_arr_expression is exited.
func (s *BaseDmSqlParserListener) ExitNew_arr_expression(ctx *New_arr_expressionContext) {}

// EnterArray_creation_expression is called when production array_creation_expression is entered.
func (s *BaseDmSqlParserListener) EnterArray_creation_expression(ctx *Array_creation_expressionContext) {
}

// ExitArray_creation_expression is called when production array_creation_expression is exited.
func (s *BaseDmSqlParserListener) ExitArray_creation_expression(ctx *Array_creation_expressionContext) {
}

// EnterPlsql_datatype_ex is called when production plsql_datatype_ex is entered.
func (s *BaseDmSqlParserListener) EnterPlsql_datatype_ex(ctx *Plsql_datatype_exContext) {}

// ExitPlsql_datatype_ex is called when production plsql_datatype_ex is exited.
func (s *BaseDmSqlParserListener) ExitPlsql_datatype_ex(ctx *Plsql_datatype_exContext) {}

// EnterNew_array_type is called when production new_array_type is entered.
func (s *BaseDmSqlParserListener) EnterNew_array_type(ctx *New_array_typeContext) {}

// ExitNew_array_type is called when production new_array_type is exited.
func (s *BaseDmSqlParserListener) ExitNew_array_type(ctx *New_array_typeContext) {}

// EnterOpt_array_initializer is called when production opt_array_initializer is entered.
func (s *BaseDmSqlParserListener) EnterOpt_array_initializer(ctx *Opt_array_initializerContext) {}

// ExitOpt_array_initializer is called when production opt_array_initializer is exited.
func (s *BaseDmSqlParserListener) ExitOpt_array_initializer(ctx *Opt_array_initializerContext) {}

// EnterArray_initializer is called when production array_initializer is entered.
func (s *BaseDmSqlParserListener) EnterArray_initializer(ctx *Array_initializerContext) {}

// ExitArray_initializer is called when production array_initializer is exited.
func (s *BaseDmSqlParserListener) ExitArray_initializer(ctx *Array_initializerContext) {}

// EnterVariable_initializer_list is called when production variable_initializer_list is entered.
func (s *BaseDmSqlParserListener) EnterVariable_initializer_list(ctx *Variable_initializer_listContext) {
}

// ExitVariable_initializer_list is called when production variable_initializer_list is exited.
func (s *BaseDmSqlParserListener) ExitVariable_initializer_list(ctx *Variable_initializer_listContext) {
}

// EnterVariable_initializer is called when production variable_initializer is entered.
func (s *BaseDmSqlParserListener) EnterVariable_initializer(ctx *Variable_initializerContext) {}

// ExitVariable_initializer is called when production variable_initializer is exited.
func (s *BaseDmSqlParserListener) ExitVariable_initializer(ctx *Variable_initializerContext) {}

// EnterOpt_comma is called when production opt_comma is entered.
func (s *BaseDmSqlParserListener) EnterOpt_comma(ctx *Opt_commaContext) {}

// ExitOpt_comma is called when production opt_comma is exited.
func (s *BaseDmSqlParserListener) ExitOpt_comma(ctx *Opt_commaContext) {}

// EnterSizeof_expression is called when production sizeof_expression is entered.
func (s *BaseDmSqlParserListener) EnterSizeof_expression(ctx *Sizeof_expressionContext) {}

// ExitSizeof_expression is called when production sizeof_expression is exited.
func (s *BaseDmSqlParserListener) ExitSizeof_expression(ctx *Sizeof_expressionContext) {}

// EnterElement_access is called when production element_access is entered.
func (s *BaseDmSqlParserListener) EnterElement_access(ctx *Element_accessContext) {}

// ExitElement_access is called when production element_access is exited.
func (s *BaseDmSqlParserListener) ExitElement_access(ctx *Element_accessContext) {}

// EnterDecode_case is called when production decode_case is entered.
func (s *BaseDmSqlParserListener) EnterDecode_case(ctx *Decode_caseContext) {}

// ExitDecode_case is called when production decode_case is exited.
func (s *BaseDmSqlParserListener) ExitDecode_case(ctx *Decode_caseContext) {}

// EnterElse_exp is called when production else_exp is entered.
func (s *BaseDmSqlParserListener) EnterElse_exp(ctx *Else_expContext) {}

// ExitElse_exp is called when production else_exp is exited.
func (s *BaseDmSqlParserListener) ExitElse_exp(ctx *Else_expContext) {}

// EnterBoolean_case is called when production boolean_case is entered.
func (s *BaseDmSqlParserListener) EnterBoolean_case(ctx *Boolean_caseContext) {}

// ExitBoolean_case is called when production boolean_case is exited.
func (s *BaseDmSqlParserListener) ExitBoolean_case(ctx *Boolean_caseContext) {}

// EnterIf_exp is called when production if_exp is entered.
func (s *BaseDmSqlParserListener) EnterIf_exp(ctx *If_expContext) {}

// ExitIf_exp is called when production if_exp is exited.
func (s *BaseDmSqlParserListener) ExitIf_exp(ctx *If_expContext) {}

// EnterBool_when_list is called when production bool_when_list is entered.
func (s *BaseDmSqlParserListener) EnterBool_when_list(ctx *Bool_when_listContext) {}

// ExitBool_when_list is called when production bool_when_list is exited.
func (s *BaseDmSqlParserListener) ExitBool_when_list(ctx *Bool_when_listContext) {}

// EnterOps is called when production ops is entered.
func (s *BaseDmSqlParserListener) EnterOps(ctx *OpsContext) {}

// ExitOps is called when production ops is exited.
func (s *BaseDmSqlParserListener) ExitOps(ctx *OpsContext) {}

// EnterValue_list is called when production value_list is entered.
func (s *BaseDmSqlParserListener) EnterValue_list(ctx *Value_listContext) {}

// ExitValue_list is called when production value_list is exited.
func (s *BaseDmSqlParserListener) ExitValue_list(ctx *Value_listContext) {}

// EnterIn_value_list is called when production in_value_list is entered.
func (s *BaseDmSqlParserListener) EnterIn_value_list(ctx *In_value_listContext) {}

// ExitIn_value_list is called when production in_value_list is exited.
func (s *BaseDmSqlParserListener) ExitIn_value_list(ctx *In_value_listContext) {}

// EnterValue_list_set is called when production value_list_set is entered.
func (s *BaseDmSqlParserListener) EnterValue_list_set(ctx *Value_list_setContext) {}

// ExitValue_list_set is called when production value_list_set is exited.
func (s *BaseDmSqlParserListener) ExitValue_list_set(ctx *Value_list_setContext) {}

// EnterComma_list is called when production comma_list is entered.
func (s *BaseDmSqlParserListener) EnterComma_list(ctx *Comma_listContext) {}

// ExitComma_list is called when production comma_list is exited.
func (s *BaseDmSqlParserListener) ExitComma_list(ctx *Comma_listContext) {}

// EnterIns_value_list is called when production ins_value_list is entered.
func (s *BaseDmSqlParserListener) EnterIns_value_list(ctx *Ins_value_listContext) {}

// ExitIns_value_list is called when production ins_value_list is exited.
func (s *BaseDmSqlParserListener) ExitIns_value_list(ctx *Ins_value_listContext) {}

// EnterNull_value is called when production null_value is entered.
func (s *BaseDmSqlParserListener) EnterNull_value(ctx *Null_valueContext) {}

// ExitNull_value is called when production null_value is exited.
func (s *BaseDmSqlParserListener) ExitNull_value(ctx *Null_valueContext) {}

// EnterId_and_rsvd_word_others is called when production id_and_rsvd_word_others is entered.
func (s *BaseDmSqlParserListener) EnterId_and_rsvd_word_others(ctx *Id_and_rsvd_word_othersContext) {}

// ExitId_and_rsvd_word_others is called when production id_and_rsvd_word_others is exited.
func (s *BaseDmSqlParserListener) ExitId_and_rsvd_word_others(ctx *Id_and_rsvd_word_othersContext) {}

// EnterId_and_rsvd_word is called when production id_and_rsvd_word is entered.
func (s *BaseDmSqlParserListener) EnterId_and_rsvd_word(ctx *Id_and_rsvd_wordContext) {}

// ExitId_and_rsvd_word is called when production id_and_rsvd_word is exited.
func (s *BaseDmSqlParserListener) ExitId_and_rsvd_word(ctx *Id_and_rsvd_wordContext) {}

// EnterStm_param is called when production stm_param is entered.
func (s *BaseDmSqlParserListener) EnterStm_param(ctx *Stm_paramContext) {}

// ExitStm_param is called when production stm_param is exited.
func (s *BaseDmSqlParserListener) ExitStm_param(ctx *Stm_paramContext) {}

// EnterStm_param_normal is called when production stm_param_normal is entered.
func (s *BaseDmSqlParserListener) EnterStm_param_normal(ctx *Stm_param_normalContext) {}

// ExitStm_param_normal is called when production stm_param_normal is exited.
func (s *BaseDmSqlParserListener) ExitStm_param_normal(ctx *Stm_param_normalContext) {}

// EnterStm_param_name is called when production stm_param_name is entered.
func (s *BaseDmSqlParserListener) EnterStm_param_name(ctx *Stm_param_nameContext) {}

// ExitStm_param_name is called when production stm_param_name is exited.
func (s *BaseDmSqlParserListener) ExitStm_param_name(ctx *Stm_param_nameContext) {}

// EnterParam_name_options is called when production param_name_options is entered.
func (s *BaseDmSqlParserListener) EnterParam_name_options(ctx *Param_name_optionsContext) {}

// ExitParam_name_options is called when production param_name_options is exited.
func (s *BaseDmSqlParserListener) ExitParam_name_options(ctx *Param_name_optionsContext) {}

// EnterContains_query_exp is called when production contains_query_exp is entered.
func (s *BaseDmSqlParserListener) EnterContains_query_exp(ctx *Contains_query_expContext) {}

// ExitContains_query_exp is called when production contains_query_exp is exited.
func (s *BaseDmSqlParserListener) ExitContains_query_exp(ctx *Contains_query_expContext) {}

// EnterContains_query_exp_lst is called when production contains_query_exp_lst is entered.
func (s *BaseDmSqlParserListener) EnterContains_query_exp_lst(ctx *Contains_query_exp_lstContext) {}

// ExitContains_query_exp_lst is called when production contains_query_exp_lst is exited.
func (s *BaseDmSqlParserListener) ExitContains_query_exp_lst(ctx *Contains_query_exp_lstContext) {}

// EnterContains_exp is called when production contains_exp is entered.
func (s *BaseDmSqlParserListener) EnterContains_exp(ctx *Contains_expContext) {}

// ExitContains_exp is called when production contains_exp is exited.
func (s *BaseDmSqlParserListener) ExitContains_exp(ctx *Contains_expContext) {}

// EnterStrict_option is called when production strict_option is entered.
func (s *BaseDmSqlParserListener) EnterStrict_option(ctx *Strict_optionContext) {}

// ExitStrict_option is called when production strict_option is exited.
func (s *BaseDmSqlParserListener) ExitStrict_option(ctx *Strict_optionContext) {}

// EnterWith_unique_option is called when production with_unique_option is entered.
func (s *BaseDmSqlParserListener) EnterWith_unique_option(ctx *With_unique_optionContext) {}

// ExitWith_unique_option is called when production with_unique_option is exited.
func (s *BaseDmSqlParserListener) ExitWith_unique_option(ctx *With_unique_optionContext) {}

// EnterType_option is called when production type_option is entered.
func (s *BaseDmSqlParserListener) EnterType_option(ctx *Type_optionContext) {}

// ExitType_option is called when production type_option is exited.
func (s *BaseDmSqlParserListener) ExitType_option(ctx *Type_optionContext) {}

// EnterType_element is called when production type_element is entered.
func (s *BaseDmSqlParserListener) EnterType_element(ctx *Type_elementContext) {}

// ExitType_element is called when production type_element is exited.
func (s *BaseDmSqlParserListener) ExitType_element(ctx *Type_elementContext) {}

// EnterType_element_list is called when production type_element_list is entered.
func (s *BaseDmSqlParserListener) EnterType_element_list(ctx *Type_element_listContext) {}

// ExitType_element_list is called when production type_element_list is exited.
func (s *BaseDmSqlParserListener) ExitType_element_list(ctx *Type_element_listContext) {}

// EnterBool_exp is called when production bool_exp is entered.
func (s *BaseDmSqlParserListener) EnterBool_exp(ctx *Bool_expContext) {}

// ExitBool_exp is called when production bool_exp is exited.
func (s *BaseDmSqlParserListener) ExitBool_exp(ctx *Bool_expContext) {}

// EnterBool_exp_element is called when production bool_exp_element is entered.
func (s *BaseDmSqlParserListener) EnterBool_exp_element(ctx *Bool_exp_elementContext) {}

// ExitBool_exp_element is called when production bool_exp_element is exited.
func (s *BaseDmSqlParserListener) ExitBool_exp_element(ctx *Bool_exp_elementContext) {}

// EnterQuery_any_options is called when production query_any_options is entered.
func (s *BaseDmSqlParserListener) EnterQuery_any_options(ctx *Query_any_optionsContext) {}

// ExitQuery_any_options is called when production query_any_options is exited.
func (s *BaseDmSqlParserListener) ExitQuery_any_options(ctx *Query_any_optionsContext) {}

// EnterGlobal_var is called when production global_var is entered.
func (s *BaseDmSqlParserListener) EnterGlobal_var(ctx *Global_varContext) {}

// ExitGlobal_var is called when production global_var is exited.
func (s *BaseDmSqlParserListener) ExitGlobal_var(ctx *Global_varContext) {}

// EnterReserved_word is called when production reserved_word is entered.
func (s *BaseDmSqlParserListener) EnterReserved_word(ctx *Reserved_wordContext) {}

// ExitReserved_word is called when production reserved_word is exited.
func (s *BaseDmSqlParserListener) ExitReserved_word(ctx *Reserved_wordContext) {}

// EnterNew_none_reserved_word is called when production new_none_reserved_word is entered.
func (s *BaseDmSqlParserListener) EnterNew_none_reserved_word(ctx *New_none_reserved_wordContext) {}

// ExitNew_none_reserved_word is called when production new_none_reserved_word is exited.
func (s *BaseDmSqlParserListener) ExitNew_none_reserved_word(ctx *New_none_reserved_wordContext) {}

// EnterInterval_nresvd_word is called when production interval_nresvd_word is entered.
func (s *BaseDmSqlParserListener) EnterInterval_nresvd_word(ctx *Interval_nresvd_wordContext) {}

// ExitInterval_nresvd_word is called when production interval_nresvd_word is exited.
func (s *BaseDmSqlParserListener) ExitInterval_nresvd_word(ctx *Interval_nresvd_wordContext) {}

// EnterVariable_resvd_word is called when production variable_resvd_word is entered.
func (s *BaseDmSqlParserListener) EnterVariable_resvd_word(ctx *Variable_resvd_wordContext) {}

// ExitVariable_resvd_word is called when production variable_resvd_word is exited.
func (s *BaseDmSqlParserListener) ExitVariable_resvd_word(ctx *Variable_resvd_wordContext) {}

// EnterAlias_resvd_word is called when production alias_resvd_word is entered.
func (s *BaseDmSqlParserListener) EnterAlias_resvd_word(ctx *Alias_resvd_wordContext) {}

// ExitAlias_resvd_word is called when production alias_resvd_word is exited.
func (s *BaseDmSqlParserListener) ExitAlias_resvd_word(ctx *Alias_resvd_wordContext) {}

// EnterSchname_resvd_word is called when production schname_resvd_word is entered.
func (s *BaseDmSqlParserListener) EnterSchname_resvd_word(ctx *Schname_resvd_wordContext) {}

// ExitSchname_resvd_word is called when production schname_resvd_word is exited.
func (s *BaseDmSqlParserListener) ExitSchname_resvd_word(ctx *Schname_resvd_wordContext) {}

// EnterRaw_id is called when production raw_id is entered.
func (s *BaseDmSqlParserListener) EnterRaw_id(ctx *Raw_idContext) {}

// ExitRaw_id is called when production raw_id is exited.
func (s *BaseDmSqlParserListener) ExitRaw_id(ctx *Raw_idContext) {}

// EnterId is called when production id is entered.
func (s *BaseDmSqlParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseDmSqlParserListener) ExitId(ctx *IdContext) {}

// EnterQualified_name is called when production qualified_name is entered.
func (s *BaseDmSqlParserListener) EnterQualified_name(ctx *Qualified_nameContext) {}

// ExitQualified_name is called when production qualified_name is exited.
func (s *BaseDmSqlParserListener) ExitQualified_name(ctx *Qualified_nameContext) {}

// EnterQualified_name2 is called when production qualified_name2 is entered.
func (s *BaseDmSqlParserListener) EnterQualified_name2(ctx *Qualified_name2Context) {}

// ExitQualified_name2 is called when production qualified_name2 is exited.
func (s *BaseDmSqlParserListener) ExitQualified_name2(ctx *Qualified_name2Context) {}

// EnterVariable_name is called when production variable_name is entered.
func (s *BaseDmSqlParserListener) EnterVariable_name(ctx *Variable_nameContext) {}

// ExitVariable_name is called when production variable_name is exited.
func (s *BaseDmSqlParserListener) ExitVariable_name(ctx *Variable_nameContext) {}

// EnterEnd_loop_label_null is called when production end_loop_label_null is entered.
func (s *BaseDmSqlParserListener) EnterEnd_loop_label_null(ctx *End_loop_label_nullContext) {}

// ExitEnd_loop_label_null is called when production end_loop_label_null is exited.
func (s *BaseDmSqlParserListener) ExitEnd_loop_label_null(ctx *End_loop_label_nullContext) {}

// EnterLabel_name_options is called when production label_name_options is entered.
func (s *BaseDmSqlParserListener) EnterLabel_name_options(ctx *Label_name_optionsContext) {}

// ExitLabel_name_options is called when production label_name_options is exited.
func (s *BaseDmSqlParserListener) ExitLabel_name_options(ctx *Label_name_optionsContext) {}

// EnterLabel_name is called when production label_name is entered.
func (s *BaseDmSqlParserListener) EnterLabel_name(ctx *Label_nameContext) {}

// ExitLabel_name is called when production label_name is exited.
func (s *BaseDmSqlParserListener) ExitLabel_name(ctx *Label_nameContext) {}

// EnterDatabase_name is called when production database_name is entered.
func (s *BaseDmSqlParserListener) EnterDatabase_name(ctx *Database_nameContext) {}

// ExitDatabase_name is called when production database_name is exited.
func (s *BaseDmSqlParserListener) ExitDatabase_name(ctx *Database_nameContext) {}

// EnterBackup_name is called when production backup_name is entered.
func (s *BaseDmSqlParserListener) EnterBackup_name(ctx *Backup_nameContext) {}

// ExitBackup_name is called when production backup_name is exited.
func (s *BaseDmSqlParserListener) ExitBackup_name(ctx *Backup_nameContext) {}

// EnterFull_proc_name is called when production full_proc_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_proc_name(ctx *Full_proc_nameContext) {}

// ExitFull_proc_name is called when production full_proc_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_proc_name(ctx *Full_proc_nameContext) {}

// EnterFull_proc_name2 is called when production full_proc_name2 is entered.
func (s *BaseDmSqlParserListener) EnterFull_proc_name2(ctx *Full_proc_name2Context) {}

// ExitFull_proc_name2 is called when production full_proc_name2 is exited.
func (s *BaseDmSqlParserListener) ExitFull_proc_name2(ctx *Full_proc_name2Context) {}

// EnterFull_fun_name is called when production full_fun_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_fun_name(ctx *Full_fun_nameContext) {}

// ExitFull_fun_name is called when production full_fun_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_fun_name(ctx *Full_fun_nameContext) {}

// EnterFull_table_name is called when production full_table_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_table_name(ctx *Full_table_nameContext) {}

// ExitFull_table_name is called when production full_table_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_table_name(ctx *Full_table_nameContext) {}

// EnterFull_grp_name is called when production full_grp_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_grp_name(ctx *Full_grp_nameContext) {}

// ExitFull_grp_name is called when production full_grp_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_grp_name(ctx *Full_grp_nameContext) {}

// EnterFull_table_name2 is called when production full_table_name2 is entered.
func (s *BaseDmSqlParserListener) EnterFull_table_name2(ctx *Full_table_name2Context) {}

// ExitFull_table_name2 is called when production full_table_name2 is exited.
func (s *BaseDmSqlParserListener) ExitFull_table_name2(ctx *Full_table_name2Context) {}

// EnterFull_partition_name is called when production full_partition_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_partition_name(ctx *Full_partition_nameContext) {}

// ExitFull_partition_name is called when production full_partition_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_partition_name(ctx *Full_partition_nameContext) {}

// EnterFull_schema_name is called when production full_schema_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_schema_name(ctx *Full_schema_nameContext) {}

// ExitFull_schema_name is called when production full_schema_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_schema_name(ctx *Full_schema_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseDmSqlParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseDmSqlParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseDmSqlParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseDmSqlParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterConstraint_name is called when production constraint_name is entered.
func (s *BaseDmSqlParserListener) EnterConstraint_name(ctx *Constraint_nameContext) {}

// ExitConstraint_name is called when production constraint_name is exited.
func (s *BaseDmSqlParserListener) ExitConstraint_name(ctx *Constraint_nameContext) {}

// EnterFull_trigger_name is called when production full_trigger_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_trigger_name(ctx *Full_trigger_nameContext) {}

// ExitFull_trigger_name is called when production full_trigger_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_trigger_name(ctx *Full_trigger_nameContext) {}

// EnterFull_trigger_name2 is called when production full_trigger_name2 is entered.
func (s *BaseDmSqlParserListener) EnterFull_trigger_name2(ctx *Full_trigger_name2Context) {}

// ExitFull_trigger_name2 is called when production full_trigger_name2 is exited.
func (s *BaseDmSqlParserListener) ExitFull_trigger_name2(ctx *Full_trigger_name2Context) {}

// EnterFull_view_name is called when production full_view_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_view_name(ctx *Full_view_nameContext) {}

// ExitFull_view_name is called when production full_view_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_view_name(ctx *Full_view_nameContext) {}

// EnterFull_view_name2 is called when production full_view_name2 is entered.
func (s *BaseDmSqlParserListener) EnterFull_view_name2(ctx *Full_view_name2Context) {}

// ExitFull_view_name2 is called when production full_view_name2 is exited.
func (s *BaseDmSqlParserListener) ExitFull_view_name2(ctx *Full_view_name2Context) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BaseDmSqlParserListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BaseDmSqlParserListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterTrigger_name is called when production trigger_name is entered.
func (s *BaseDmSqlParserListener) EnterTrigger_name(ctx *Trigger_nameContext) {}

// ExitTrigger_name is called when production trigger_name is exited.
func (s *BaseDmSqlParserListener) ExitTrigger_name(ctx *Trigger_nameContext) {}

// EnterLogin_name is called when production login_name is entered.
func (s *BaseDmSqlParserListener) EnterLogin_name(ctx *Login_nameContext) {}

// ExitLogin_name is called when production login_name is exited.
func (s *BaseDmSqlParserListener) ExitLogin_name(ctx *Login_nameContext) {}

// EnterProfile_name is called when production profile_name is entered.
func (s *BaseDmSqlParserListener) EnterProfile_name(ctx *Profile_nameContext) {}

// ExitProfile_name is called when production profile_name is exited.
func (s *BaseDmSqlParserListener) ExitProfile_name(ctx *Profile_nameContext) {}

// EnterUser_name is called when production user_name is entered.
func (s *BaseDmSqlParserListener) EnterUser_name(ctx *User_nameContext) {}

// ExitUser_name is called when production user_name is exited.
func (s *BaseDmSqlParserListener) ExitUser_name(ctx *User_nameContext) {}

// EnterRole_name is called when production role_name is entered.
func (s *BaseDmSqlParserListener) EnterRole_name(ctx *Role_nameContext) {}

// ExitRole_name is called when production role_name is exited.
func (s *BaseDmSqlParserListener) ExitRole_name(ctx *Role_nameContext) {}

// EnterUser_role_name is called when production user_role_name is entered.
func (s *BaseDmSqlParserListener) EnterUser_role_name(ctx *User_role_nameContext) {}

// ExitUser_role_name is called when production user_role_name is exited.
func (s *BaseDmSqlParserListener) ExitUser_role_name(ctx *User_role_nameContext) {}

// EnterRole_name_list is called when production role_name_list is entered.
func (s *BaseDmSqlParserListener) EnterRole_name_list(ctx *Role_name_listContext) {}

// ExitRole_name_list is called when production role_name_list is exited.
func (s *BaseDmSqlParserListener) ExitRole_name_list(ctx *Role_name_listContext) {}

// EnterFull_func_name is called when production full_func_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_func_name(ctx *Full_func_nameContext) {}

// ExitFull_func_name is called when production full_func_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_func_name(ctx *Full_func_nameContext) {}

// EnterParam_name is called when production param_name is entered.
func (s *BaseDmSqlParserListener) EnterParam_name(ctx *Param_nameContext) {}

// ExitParam_name is called when production param_name is exited.
func (s *BaseDmSqlParserListener) ExitParam_name(ctx *Param_nameContext) {}

// EnterIndex_name is called when production index_name is entered.
func (s *BaseDmSqlParserListener) EnterIndex_name(ctx *Index_nameContext) {}

// ExitIndex_name is called when production index_name is exited.
func (s *BaseDmSqlParserListener) ExitIndex_name(ctx *Index_nameContext) {}

// EnterIndex_name2 is called when production index_name2 is entered.
func (s *BaseDmSqlParserListener) EnterIndex_name2(ctx *Index_name2Context) {}

// ExitIndex_name2 is called when production index_name2 is exited.
func (s *BaseDmSqlParserListener) ExitIndex_name2(ctx *Index_name2Context) {}

// EnterTrig_old_name is called when production trig_old_name is entered.
func (s *BaseDmSqlParserListener) EnterTrig_old_name(ctx *Trig_old_nameContext) {}

// ExitTrig_old_name is called when production trig_old_name is exited.
func (s *BaseDmSqlParserListener) ExitTrig_old_name(ctx *Trig_old_nameContext) {}

// EnterTrig_new_name is called when production trig_new_name is entered.
func (s *BaseDmSqlParserListener) EnterTrig_new_name(ctx *Trig_new_nameContext) {}

// ExitTrig_new_name is called when production trig_new_name is exited.
func (s *BaseDmSqlParserListener) ExitTrig_new_name(ctx *Trig_new_nameContext) {}

// EnterFull_tv_name is called when production full_tv_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_tv_name(ctx *Full_tv_nameContext) {}

// ExitFull_tv_name is called when production full_tv_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_tv_name(ctx *Full_tv_nameContext) {}

// EnterFull_object_name is called when production full_object_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_object_name(ctx *Full_object_nameContext) {}

// ExitFull_object_name is called when production full_object_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_object_name(ctx *Full_object_nameContext) {}

// EnterOrient_option is called when production orient_option is entered.
func (s *BaseDmSqlParserListener) EnterOrient_option(ctx *Orient_optionContext) {}

// ExitOrient_option is called when production orient_option is exited.
func (s *BaseDmSqlParserListener) ExitOrient_option(ctx *Orient_optionContext) {}

// EnterDatepart is called when production datepart is entered.
func (s *BaseDmSqlParserListener) EnterDatepart(ctx *DatepartContext) {}

// ExitDatepart is called when production datepart is exited.
func (s *BaseDmSqlParserListener) ExitDatepart(ctx *DatepartContext) {}

// EnterDatepart_op is called when production datepart_op is entered.
func (s *BaseDmSqlParserListener) EnterDatepart_op(ctx *Datepart_opContext) {}

// ExitDatepart_op is called when production datepart_op is exited.
func (s *BaseDmSqlParserListener) ExitDatepart_op(ctx *Datepart_opContext) {}

// EnterDatead_fun is called when production datead_fun is entered.
func (s *BaseDmSqlParserListener) EnterDatead_fun(ctx *Datead_funContext) {}

// ExitDatead_fun is called when production datead_fun is exited.
func (s *BaseDmSqlParserListener) ExitDatead_fun(ctx *Datead_funContext) {}

// EnterReturning is called when production returning is entered.
func (s *BaseDmSqlParserListener) EnterReturning(ctx *ReturningContext) {}

// ExitReturning is called when production returning is exited.
func (s *BaseDmSqlParserListener) ExitReturning(ctx *ReturningContext) {}

// EnterPretty is called when production pretty is entered.
func (s *BaseDmSqlParserListener) EnterPretty(ctx *PrettyContext) {}

// ExitPretty is called when production pretty is exited.
func (s *BaseDmSqlParserListener) ExitPretty(ctx *PrettyContext) {}

// EnterWrapper_flag is called when production wrapper_flag is entered.
func (s *BaseDmSqlParserListener) EnterWrapper_flag(ctx *Wrapper_flagContext) {}

// ExitWrapper_flag is called when production wrapper_flag is exited.
func (s *BaseDmSqlParserListener) ExitWrapper_flag(ctx *Wrapper_flagContext) {}

// EnterArray_wrapper is called when production array_wrapper is entered.
func (s *BaseDmSqlParserListener) EnterArray_wrapper(ctx *Array_wrapperContext) {}

// ExitArray_wrapper is called when production array_wrapper is exited.
func (s *BaseDmSqlParserListener) ExitArray_wrapper(ctx *Array_wrapperContext) {}

// EnterJson_tail_on_empty is called when production json_tail_on_empty is entered.
func (s *BaseDmSqlParserListener) EnterJson_tail_on_empty(ctx *Json_tail_on_emptyContext) {}

// ExitJson_tail_on_empty is called when production json_tail_on_empty is exited.
func (s *BaseDmSqlParserListener) ExitJson_tail_on_empty(ctx *Json_tail_on_emptyContext) {}

// EnterEmpty_handle is called when production empty_handle is entered.
func (s *BaseDmSqlParserListener) EnterEmpty_handle(ctx *Empty_handleContext) {}

// ExitEmpty_handle is called when production empty_handle is exited.
func (s *BaseDmSqlParserListener) ExitEmpty_handle(ctx *Empty_handleContext) {}

// EnterJson_tail_on_error_null is called when production json_tail_on_error_null is entered.
func (s *BaseDmSqlParserListener) EnterJson_tail_on_error_null(ctx *Json_tail_on_error_nullContext) {}

// ExitJson_tail_on_error_null is called when production json_tail_on_error_null is exited.
func (s *BaseDmSqlParserListener) ExitJson_tail_on_error_null(ctx *Json_tail_on_error_nullContext) {}

// EnterError_handle is called when production error_handle is entered.
func (s *BaseDmSqlParserListener) EnterError_handle(ctx *Error_handleContext) {}

// ExitError_handle is called when production error_handle is exited.
func (s *BaseDmSqlParserListener) ExitError_handle(ctx *Error_handleContext) {}

// EnterSavepoint_name is called when production savepoint_name is entered.
func (s *BaseDmSqlParserListener) EnterSavepoint_name(ctx *Savepoint_nameContext) {}

// ExitSavepoint_name is called when production savepoint_name is exited.
func (s *BaseDmSqlParserListener) ExitSavepoint_name(ctx *Savepoint_nameContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseDmSqlParserListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseDmSqlParserListener) ExitAlias(ctx *AliasContext) {}

// EnterAlias_2 is called when production alias_2 is entered.
func (s *BaseDmSqlParserListener) EnterAlias_2(ctx *Alias_2Context) {}

// ExitAlias_2 is called when production alias_2 is exited.
func (s *BaseDmSqlParserListener) ExitAlias_2(ctx *Alias_2Context) {}

// EnterFull_column_name is called when production full_column_name is entered.
func (s *BaseDmSqlParserListener) EnterFull_column_name(ctx *Full_column_nameContext) {}

// ExitFull_column_name is called when production full_column_name is exited.
func (s *BaseDmSqlParserListener) ExitFull_column_name(ctx *Full_column_nameContext) {}

// EnterSchema_name is called when production schema_name is entered.
func (s *BaseDmSqlParserListener) EnterSchema_name(ctx *Schema_nameContext) {}

// ExitSchema_name is called when production schema_name is exited.
func (s *BaseDmSqlParserListener) ExitSchema_name(ctx *Schema_nameContext) {}

// EnterNot_tag is called when production not_tag is entered.
func (s *BaseDmSqlParserListener) EnterNot_tag(ctx *Not_tagContext) {}

// ExitNot_tag is called when production not_tag is exited.
func (s *BaseDmSqlParserListener) ExitNot_tag(ctx *Not_tagContext) {}

// EnterDebug_tag is called when production debug_tag is entered.
func (s *BaseDmSqlParserListener) EnterDebug_tag(ctx *Debug_tagContext) {}

// ExitDebug_tag is called when production debug_tag is exited.
func (s *BaseDmSqlParserListener) ExitDebug_tag(ctx *Debug_tagContext) {}

// EnterColumn_tag is called when production column_tag is entered.
func (s *BaseDmSqlParserListener) EnterColumn_tag(ctx *Column_tagContext) {}

// ExitColumn_tag is called when production column_tag is exited.
func (s *BaseDmSqlParserListener) ExitColumn_tag(ctx *Column_tagContext) {}

// EnterPendant_tag is called when production pendant_tag is entered.
func (s *BaseDmSqlParserListener) EnterPendant_tag(ctx *Pendant_tagContext) {}

// ExitPendant_tag is called when production pendant_tag is exited.
func (s *BaseDmSqlParserListener) ExitPendant_tag(ctx *Pendant_tagContext) {}

// EnterUnique_tag is called when production unique_tag is entered.
func (s *BaseDmSqlParserListener) EnterUnique_tag(ctx *Unique_tagContext) {}

// ExitUnique_tag is called when production unique_tag is exited.
func (s *BaseDmSqlParserListener) ExitUnique_tag(ctx *Unique_tagContext) {}

// EnterPartition_tag is called when production partition_tag is entered.
func (s *BaseDmSqlParserListener) EnterPartition_tag(ctx *Partition_tagContext) {}

// ExitPartition_tag is called when production partition_tag is exited.
func (s *BaseDmSqlParserListener) ExitPartition_tag(ctx *Partition_tagContext) {}

// EnterRow_tag is called when production row_tag is entered.
func (s *BaseDmSqlParserListener) EnterRow_tag(ctx *Row_tagContext) {}

// ExitRow_tag is called when production row_tag is exited.
func (s *BaseDmSqlParserListener) ExitRow_tag(ctx *Row_tagContext) {}

// EnterAs_tag is called when production as_tag is entered.
func (s *BaseDmSqlParserListener) EnterAs_tag(ctx *As_tagContext) {}

// ExitAs_tag is called when production as_tag is exited.
func (s *BaseDmSqlParserListener) ExitAs_tag(ctx *As_tagContext) {}

// EnterFrom_tag is called when production from_tag is entered.
func (s *BaseDmSqlParserListener) EnterFrom_tag(ctx *From_tagContext) {}

// ExitFrom_tag is called when production from_tag is exited.
func (s *BaseDmSqlParserListener) ExitFrom_tag(ctx *From_tagContext) {}

// EnterInto_tag is called when production into_tag is entered.
func (s *BaseDmSqlParserListener) EnterInto_tag(ctx *Into_tagContext) {}

// ExitInto_tag is called when production into_tag is exited.
func (s *BaseDmSqlParserListener) ExitInto_tag(ctx *Into_tagContext) {}

// EnterWork_tag is called when production work_tag is entered.
func (s *BaseDmSqlParserListener) EnterWork_tag(ctx *Work_tagContext) {}

// ExitWork_tag is called when production work_tag is exited.
func (s *BaseDmSqlParserListener) ExitWork_tag(ctx *Work_tagContext) {}

// EnterWith_grant_option is called when production with_grant_option is entered.
func (s *BaseDmSqlParserListener) EnterWith_grant_option(ctx *With_grant_optionContext) {}

// ExitWith_grant_option is called when production with_grant_option is exited.
func (s *BaseDmSqlParserListener) ExitWith_grant_option(ctx *With_grant_optionContext) {}

// EnterWith_admin_option is called when production with_admin_option is entered.
func (s *BaseDmSqlParserListener) EnterWith_admin_option(ctx *With_admin_optionContext) {}

// ExitWith_admin_option is called when production with_admin_option is exited.
func (s *BaseDmSqlParserListener) ExitWith_admin_option(ctx *With_admin_optionContext) {}

// EnterTime_zone_or_local is called when production time_zone_or_local is entered.
func (s *BaseDmSqlParserListener) EnterTime_zone_or_local(ctx *Time_zone_or_localContext) {}

// ExitTime_zone_or_local is called when production time_zone_or_local is exited.
func (s *BaseDmSqlParserListener) ExitTime_zone_or_local(ctx *Time_zone_or_localContext) {}

// EnterSub_plsql_datatype is called when production sub_plsql_datatype is entered.
func (s *BaseDmSqlParserListener) EnterSub_plsql_datatype(ctx *Sub_plsql_datatypeContext) {}

// ExitSub_plsql_datatype is called when production sub_plsql_datatype is exited.
func (s *BaseDmSqlParserListener) ExitSub_plsql_datatype(ctx *Sub_plsql_datatypeContext) {}

// EnterDatatype_list is called when production datatype_list is entered.
func (s *BaseDmSqlParserListener) EnterDatatype_list(ctx *Datatype_listContext) {}

// ExitDatatype_list is called when production datatype_list is exited.
func (s *BaseDmSqlParserListener) ExitDatatype_list(ctx *Datatype_listContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseDmSqlParserListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseDmSqlParserListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterDatatype2 is called when production datatype2 is entered.
func (s *BaseDmSqlParserListener) EnterDatatype2(ctx *Datatype2Context) {}

// ExitDatatype2 is called when production datatype2 is exited.
func (s *BaseDmSqlParserListener) ExitDatatype2(ctx *Datatype2Context) {}

// EnterOpr_dtype is called when production opr_dtype is entered.
func (s *BaseDmSqlParserListener) EnterOpr_dtype(ctx *Opr_dtypeContext) {}

// ExitOpr_dtype is called when production opr_dtype is exited.
func (s *BaseDmSqlParserListener) ExitOpr_dtype(ctx *Opr_dtypeContext) {}

// EnterOpr_datatype_lst is called when production opr_datatype_lst is entered.
func (s *BaseDmSqlParserListener) EnterOpr_datatype_lst(ctx *Opr_datatype_lstContext) {}

// ExitOpr_datatype_lst is called when production opr_datatype_lst is exited.
func (s *BaseDmSqlParserListener) ExitOpr_datatype_lst(ctx *Opr_datatype_lstContext) {}

// EnterInterval_qualifier is called when production interval_qualifier is entered.
func (s *BaseDmSqlParserListener) EnterInterval_qualifier(ctx *Interval_qualifierContext) {}

// ExitInterval_qualifier is called when production interval_qualifier is exited.
func (s *BaseDmSqlParserListener) ExitInterval_qualifier(ctx *Interval_qualifierContext) {}

// EnterDtype is called when production dtype is entered.
func (s *BaseDmSqlParserListener) EnterDtype(ctx *DtypeContext) {}

// ExitDtype is called when production dtype is exited.
func (s *BaseDmSqlParserListener) ExitDtype(ctx *DtypeContext) {}

// EnterDtype1 is called when production dtype1 is entered.
func (s *BaseDmSqlParserListener) EnterDtype1(ctx *Dtype1Context) {}

// ExitDtype1 is called when production dtype1 is exited.
func (s *BaseDmSqlParserListener) ExitDtype1(ctx *Dtype1Context) {}

// EnterDtype2 is called when production dtype2 is entered.
func (s *BaseDmSqlParserListener) EnterDtype2(ctx *Dtype2Context) {}

// ExitDtype2 is called when production dtype2 is exited.
func (s *BaseDmSqlParserListener) ExitDtype2(ctx *Dtype2Context) {}

// EnterDouble_length_option is called when production double_length_option is entered.
func (s *BaseDmSqlParserListener) EnterDouble_length_option(ctx *Double_length_optionContext) {}

// ExitDouble_length_option is called when production double_length_option is exited.
func (s *BaseDmSqlParserListener) ExitDouble_length_option(ctx *Double_length_optionContext) {}

// EnterSize_unit_caluse is called when production size_unit_caluse is entered.
func (s *BaseDmSqlParserListener) EnterSize_unit_caluse(ctx *Size_unit_caluseContext) {}

// ExitSize_unit_caluse is called when production size_unit_caluse is exited.
func (s *BaseDmSqlParserListener) ExitSize_unit_caluse(ctx *Size_unit_caluseContext) {}

// EnterLt_integer_negative is called when production lt_integer_negative is entered.
func (s *BaseDmSqlParserListener) EnterLt_integer_negative(ctx *Lt_integer_negativeContext) {}

// ExitLt_integer_negative is called when production lt_integer_negative is exited.
func (s *BaseDmSqlParserListener) ExitLt_integer_negative(ctx *Lt_integer_negativeContext) {}

// EnterCreate_contextindex_stmt is called when production create_contextindex_stmt is entered.
func (s *BaseDmSqlParserListener) EnterCreate_contextindex_stmt(ctx *Create_contextindex_stmtContext) {
}

// ExitCreate_contextindex_stmt is called when production create_contextindex_stmt is exited.
func (s *BaseDmSqlParserListener) ExitCreate_contextindex_stmt(ctx *Create_contextindex_stmtContext) {
}

// EnterLexer_name is called when production lexer_name is entered.
func (s *BaseDmSqlParserListener) EnterLexer_name(ctx *Lexer_nameContext) {}

// ExitLexer_name is called when production lexer_name is exited.
func (s *BaseDmSqlParserListener) ExitLexer_name(ctx *Lexer_nameContext) {}

// EnterLexer_clause is called when production lexer_clause is entered.
func (s *BaseDmSqlParserListener) EnterLexer_clause(ctx *Lexer_clauseContext) {}

// ExitLexer_clause is called when production lexer_clause is exited.
func (s *BaseDmSqlParserListener) ExitLexer_clause(ctx *Lexer_clauseContext) {}

// EnterLexer_clause2 is called when production lexer_clause2 is entered.
func (s *BaseDmSqlParserListener) EnterLexer_clause2(ctx *Lexer_clause2Context) {}

// ExitLexer_clause2 is called when production lexer_clause2 is exited.
func (s *BaseDmSqlParserListener) ExitLexer_clause2(ctx *Lexer_clause2Context) {}

// EnterSync is called when production sync is entered.
func (s *BaseDmSqlParserListener) EnterSync(ctx *SyncContext) {}

// ExitSync is called when production sync is exited.
func (s *BaseDmSqlParserListener) ExitSync(ctx *SyncContext) {}

// EnterDrop_contextindex_stmt is called when production drop_contextindex_stmt is entered.
func (s *BaseDmSqlParserListener) EnterDrop_contextindex_stmt(ctx *Drop_contextindex_stmtContext) {}

// ExitDrop_contextindex_stmt is called when production drop_contextindex_stmt is exited.
func (s *BaseDmSqlParserListener) ExitDrop_contextindex_stmt(ctx *Drop_contextindex_stmtContext) {}

// EnterAlter_contextindex_stmt is called when production alter_contextindex_stmt is entered.
func (s *BaseDmSqlParserListener) EnterAlter_contextindex_stmt(ctx *Alter_contextindex_stmtContext) {}

// ExitAlter_contextindex_stmt is called when production alter_contextindex_stmt is exited.
func (s *BaseDmSqlParserListener) ExitAlter_contextindex_stmt(ctx *Alter_contextindex_stmtContext) {}

// EnterCti_sync_option is called when production cti_sync_option is entered.
func (s *BaseDmSqlParserListener) EnterCti_sync_option(ctx *Cti_sync_optionContext) {}

// ExitCti_sync_option is called when production cti_sync_option is exited.
func (s *BaseDmSqlParserListener) ExitCti_sync_option(ctx *Cti_sync_optionContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseDmSqlParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseDmSqlParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterSizeof_type is called when production sizeof_type is entered.
func (s *BaseDmSqlParserListener) EnterSizeof_type(ctx *Sizeof_typeContext) {}

// ExitSizeof_type is called when production sizeof_type is exited.
func (s *BaseDmSqlParserListener) ExitSizeof_type(ctx *Sizeof_typeContext) {}

// EnterType is called when production type is entered.
func (s *BaseDmSqlParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseDmSqlParserListener) ExitType(ctx *TypeContext) {}

// EnterArray_type is called when production array_type is entered.
func (s *BaseDmSqlParserListener) EnterArray_type(ctx *Array_typeContext) {}

// ExitArray_type is called when production array_type is exited.
func (s *BaseDmSqlParserListener) ExitArray_type(ctx *Array_typeContext) {}

// EnterBuiltin_types is called when production builtin_types is entered.
func (s *BaseDmSqlParserListener) EnterBuiltin_types(ctx *Builtin_typesContext) {}

// ExitBuiltin_types is called when production builtin_types is exited.
func (s *BaseDmSqlParserListener) ExitBuiltin_types(ctx *Builtin_typesContext) {}

// EnterIntegral_type is called when production integral_type is entered.
func (s *BaseDmSqlParserListener) EnterIntegral_type(ctx *Integral_typeContext) {}

// ExitIntegral_type is called when production integral_type is exited.
func (s *BaseDmSqlParserListener) ExitIntegral_type(ctx *Integral_typeContext) {}

// EnterSql_builtin_types is called when production sql_builtin_types is entered.
func (s *BaseDmSqlParserListener) EnterSql_builtin_types(ctx *Sql_builtin_typesContext) {}

// ExitSql_builtin_types is called when production sql_builtin_types is exited.
func (s *BaseDmSqlParserListener) ExitSql_builtin_types(ctx *Sql_builtin_typesContext) {}

// EnterCursor_declaration is called when production cursor_declaration is entered.
func (s *BaseDmSqlParserListener) EnterCursor_declaration(ctx *Cursor_declarationContext) {}

// ExitCursor_declaration is called when production cursor_declaration is exited.
func (s *BaseDmSqlParserListener) ExitCursor_declaration(ctx *Cursor_declarationContext) {}

// EnterCursor_declaration_2 is called when production cursor_declaration_2 is entered.
func (s *BaseDmSqlParserListener) EnterCursor_declaration_2(ctx *Cursor_declaration_2Context) {}

// ExitCursor_declaration_2 is called when production cursor_declaration_2 is exited.
func (s *BaseDmSqlParserListener) ExitCursor_declaration_2(ctx *Cursor_declaration_2Context) {}

// EnterCursor_attrs_options is called when production cursor_attrs_options is entered.
func (s *BaseDmSqlParserListener) EnterCursor_attrs_options(ctx *Cursor_attrs_optionsContext) {}

// ExitCursor_attrs_options is called when production cursor_attrs_options is exited.
func (s *BaseDmSqlParserListener) ExitCursor_attrs_options(ctx *Cursor_attrs_optionsContext) {}

// EnterCursor_attrs is called when production cursor_attrs is entered.
func (s *BaseDmSqlParserListener) EnterCursor_attrs(ctx *Cursor_attrsContext) {}

// ExitCursor_attrs is called when production cursor_attrs is exited.
func (s *BaseDmSqlParserListener) ExitCursor_attrs(ctx *Cursor_attrsContext) {}

// EnterCursor_attr is called when production cursor_attr is entered.
func (s *BaseDmSqlParserListener) EnterCursor_attr(ctx *Cursor_attrContext) {}

// ExitCursor_attr is called when production cursor_attr is exited.
func (s *BaseDmSqlParserListener) ExitCursor_attr(ctx *Cursor_attrContext) {}

// EnterOpt_rank_specifier is called when production opt_rank_specifier is entered.
func (s *BaseDmSqlParserListener) EnterOpt_rank_specifier(ctx *Opt_rank_specifierContext) {}

// ExitOpt_rank_specifier is called when production opt_rank_specifier is exited.
func (s *BaseDmSqlParserListener) ExitOpt_rank_specifier(ctx *Opt_rank_specifierContext) {}

// EnterRank_specifiers is called when production rank_specifiers is entered.
func (s *BaseDmSqlParserListener) EnterRank_specifiers(ctx *Rank_specifiersContext) {}

// ExitRank_specifiers is called when production rank_specifiers is exited.
func (s *BaseDmSqlParserListener) ExitRank_specifiers(ctx *Rank_specifiersContext) {}

// EnterRank_specifier is called when production rank_specifier is entered.
func (s *BaseDmSqlParserListener) EnterRank_specifier(ctx *Rank_specifierContext) {}

// ExitRank_specifier is called when production rank_specifier is exited.
func (s *BaseDmSqlParserListener) ExitRank_specifier(ctx *Rank_specifierContext) {}

// EnterOpt_dim_separators is called when production opt_dim_separators is entered.
func (s *BaseDmSqlParserListener) EnterOpt_dim_separators(ctx *Opt_dim_separatorsContext) {}

// ExitOpt_dim_separators is called when production opt_dim_separators is exited.
func (s *BaseDmSqlParserListener) ExitOpt_dim_separators(ctx *Opt_dim_separatorsContext) {}

// EnterOpt_rank_specifier2 is called when production opt_rank_specifier2 is entered.
func (s *BaseDmSqlParserListener) EnterOpt_rank_specifier2(ctx *Opt_rank_specifier2Context) {}

// ExitOpt_rank_specifier2 is called when production opt_rank_specifier2 is exited.
func (s *BaseDmSqlParserListener) ExitOpt_rank_specifier2(ctx *Opt_rank_specifier2Context) {}

// EnterDim_separators is called when production dim_separators is entered.
func (s *BaseDmSqlParserListener) EnterDim_separators(ctx *Dim_separatorsContext) {}

// ExitDim_separators is called when production dim_separators is exited.
func (s *BaseDmSqlParserListener) ExitDim_separators(ctx *Dim_separatorsContext) {}

// EnterOpt_argument_list is called when production opt_argument_list is entered.
func (s *BaseDmSqlParserListener) EnterOpt_argument_list(ctx *Opt_argument_listContext) {}

// ExitOpt_argument_list is called when production opt_argument_list is exited.
func (s *BaseDmSqlParserListener) ExitOpt_argument_list(ctx *Opt_argument_listContext) {}

// EnterJson_fun_tail is called when production json_fun_tail is entered.
func (s *BaseDmSqlParserListener) EnterJson_fun_tail(ctx *Json_fun_tailContext) {}

// ExitJson_fun_tail is called when production json_fun_tail is exited.
func (s *BaseDmSqlParserListener) ExitJson_fun_tail(ctx *Json_fun_tailContext) {}

// EnterIgnore_nulls_clause is called when production ignore_nulls_clause is entered.
func (s *BaseDmSqlParserListener) EnterIgnore_nulls_clause(ctx *Ignore_nulls_clauseContext) {}

// ExitIgnore_nulls_clause is called when production ignore_nulls_clause is exited.
func (s *BaseDmSqlParserListener) ExitIgnore_nulls_clause(ctx *Ignore_nulls_clauseContext) {}

// EnterMixed_param_list is called when production mixed_param_list is entered.
func (s *BaseDmSqlParserListener) EnterMixed_param_list(ctx *Mixed_param_listContext) {}

// ExitMixed_param_list is called when production mixed_param_list is exited.
func (s *BaseDmSqlParserListener) ExitMixed_param_list(ctx *Mixed_param_listContext) {}

// EnterMixed_param is called when production mixed_param is entered.
func (s *BaseDmSqlParserListener) EnterMixed_param(ctx *Mixed_paramContext) {}

// ExitMixed_param is called when production mixed_param is exited.
func (s *BaseDmSqlParserListener) ExitMixed_param(ctx *Mixed_paramContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseDmSqlParserListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseDmSqlParserListener) ExitArgument(ctx *ArgumentContext) {}

// EnterCursor_option is called when production cursor_option is entered.
func (s *BaseDmSqlParserListener) EnterCursor_option(ctx *Cursor_optionContext) {}

// ExitCursor_option is called when production cursor_option is exited.
func (s *BaseDmSqlParserListener) ExitCursor_option(ctx *Cursor_optionContext) {}

// EnterWithout_into_select2 is called when production without_into_select2 is entered.
func (s *BaseDmSqlParserListener) EnterWithout_into_select2(ctx *Without_into_select2Context) {}

// ExitWithout_into_select2 is called when production without_into_select2 is exited.
func (s *BaseDmSqlParserListener) ExitWithout_into_select2(ctx *Without_into_select2Context) {}

// EnterCursor_option_2 is called when production cursor_option_2 is entered.
func (s *BaseDmSqlParserListener) EnterCursor_option_2(ctx *Cursor_option_2Context) {}

// ExitCursor_option_2 is called when production cursor_option_2 is exited.
func (s *BaseDmSqlParserListener) ExitCursor_option_2(ctx *Cursor_option_2Context) {}

// EnterRegion_size is called when production region_size is entered.
func (s *BaseDmSqlParserListener) EnterRegion_size(ctx *Region_sizeContext) {}

// ExitRegion_size is called when production region_size is exited.
func (s *BaseDmSqlParserListener) ExitRegion_size(ctx *Region_sizeContext) {}

// EnterCopy_num is called when production copy_num is entered.
func (s *BaseDmSqlParserListener) EnterCopy_num(ctx *Copy_numContext) {}

// ExitCopy_num is called when production copy_num is exited.
func (s *BaseDmSqlParserListener) ExitCopy_num(ctx *Copy_numContext) {}

// EnterRedundancy_clause is called when production redundancy_clause is entered.
func (s *BaseDmSqlParserListener) EnterRedundancy_clause(ctx *Redundancy_clauseContext) {}

// ExitRedundancy_clause is called when production redundancy_clause is exited.
func (s *BaseDmSqlParserListener) ExitRedundancy_clause(ctx *Redundancy_clauseContext) {}

// EnterStriping_clause is called when production striping_clause is entered.
func (s *BaseDmSqlParserListener) EnterStriping_clause(ctx *Striping_clauseContext) {}

// ExitStriping_clause is called when production striping_clause is exited.
func (s *BaseDmSqlParserListener) ExitStriping_clause(ctx *Striping_clauseContext) {}

// EnterWith_huge_clause is called when production with_huge_clause is entered.
func (s *BaseDmSqlParserListener) EnterWith_huge_clause(ctx *With_huge_clauseContext) {}

// ExitWith_huge_clause is called when production with_huge_clause is exited.
func (s *BaseDmSqlParserListener) ExitWith_huge_clause(ctx *With_huge_clauseContext) {}
